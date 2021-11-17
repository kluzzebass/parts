package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
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

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.TenantID, input.Name)

	var dst *model.User
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListUsers(ctx context.Context, ids *[]string) ([]*model.User, error) {
	sql := `
		SELECT
			user_id,
			tenant_id,
			created_at,
			name
		FROM
			"user"
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND user_id = ANY ($1::uuid[]))
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.User
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
