package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
)

func (repo *Repository) UpsertUser(ctx context.Context, input model.NewUser) (*model.User, error) {
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
		ON CONFLICT (user_id) DO UPDATE
		SET
			name = EXCLUDED.name
		RETURNING user_id, tenant_id, created_at, name
	`

	var id string
	var tenantId string
	var createdAt time.Time
	var name string

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.TenantID, input.Name).Scan(&id, &tenantId, &createdAt, &name)

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

func (repo *Repository) ListUsers(ctx context.Context, ids *[]string) ([]*model.User, error) {
	sql := `
		SELECT
			user_id AS id,
			tenant_id,
			created_at,
			name
		FROM
			"user"
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND user_id = ANY ($1::uuid[]))
	`

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var id string
		var tenantId string
		var createdAt time.Time
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

	return users, err
}
