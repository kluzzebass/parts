package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
)

func (repo *Repository) CreateUser(nt model.NewUser) (*model.User, error) {
	sql := `
		INSERT INTO
			"user"
		(
			user_id,
			tenant_id,
			name
		)
		VALUES
		(
			COALESCE($1, gen_random_uuid()),
			$2,
			$3
		)
		RETURNING user_id, tenant_id, created_at::text, name
	`
	var id string
	var tenantId string
	var createdAt string
	var name string

	err := repo.pool.QueryRow(context.TODO(), sql, nt.ID, nt.TenantID, nt.Name).Scan(&id, &tenantId, &createdAt, &name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.User{
		ID:        id,
		TenantID:  tenantId,
		CreatedAt: createdAt,
		Name:      name,
	}, err
}

func (repo *Repository) ListUsers(userIds *[]string) ([]*model.User, error) {
	sql := `
		SELECT
			user_id AS id,
			tenant_id,
			created_at::text,
			name
		FROM
			"user"
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND user_id = ANY ($1::uuid[]))
	`
	rows, err := repo.pool.Query(context.TODO(), sql, userIds)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var id string
		var tenantId string
		var createdAt string
		var name string
		err := rows.Scan(&id, &tenantId, &createdAt, &name)
		if err != nil {
			return nil, err
		}

		users = append(users, &model.User{
			ID:        id,
			TenantID:  tenantId,
			CreatedAt: createdAt,
			Name:      name,
		})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return users, err
}
