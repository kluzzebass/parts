package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
)

func (repo *Repository) CreateTenant(nt model.NewTenant) (*model.Tenant, error) {
	sql := `
		INSERT INTO
			tenant
		(
			tenant_id,
			name
		)
		VALUES
		(
			COALESCE($1, gen_random_uuid()),
			$2
		)
		RETURNING tenant_id, created_at::text, name
	`
	var id string
	var createdAt string
	var name string

	fmt.Println("tenant_id =", *nt.ID)

	err := repo.pool.QueryRow(context.TODO(), sql, nt.ID, nt.Name).Scan(&id, &createdAt, &name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.Tenant{
		ID:        id,
		CreatedAt: createdAt,
		Name:      name,
	}, err
}

func (repo *Repository) ListTenants(id *string) ([]*model.Tenant, error) {
	sql := `
		SELECT
			t.tenant_id AS id,
			t.created_at::text,
			t.name,
			array_agg(u.user_id) as users
		FROM
			tenant t
			JOIN "user" u USING (tenant_id)
		WHERE
			$1::uuid IS NULL
			OR ($1::uuid IS NOT NULL AND t.tenant_id = $1::uuid)
		GROUP BY
			t.tenant_id
	`
	rows, err := repo.pool.Query(context.TODO(), sql, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tenants []*model.Tenant

	for rows.Next() {
		var id string
		var createdAt string
		var name string
		var users []string
		err := rows.Scan(&id, &createdAt, &name, &users)
		if err != nil {
			return nil, err
		}

		tenants = append(tenants, &model.Tenant{
			ID:        id,
			CreatedAt: createdAt,
			Name:      name,
			UserIDs:   users,
		})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return tenants, err
}
