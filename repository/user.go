package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	sql := `
		WITH
		inserted_rows AS (
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
			WHERE
				"user".name IS DISTINCT FROM EXCLUDED.name
			RETURNING
				user_id,
				tenant_id,
				created_at,
				name
		),
		selected_rows AS (
			SELECT
				user_id,
				tenant_id,
				created_at,
				name
			FROM
				"user"
			WHERE
				user_id = $1
		)
		SELECT
			COALESCE(ir.user_id, sr.user_id) AS user_id,
			COALESCE(ir.tenant_id, sr.tenant_id) AS tenant_id,
			COALESCE(ir.created_at, sr.created_at) AS created_at,
			COALESCE(ir.name, sr.name) AS name
		FROM
			inserted_rows ir
			FULL JOIN selected_rows sr USING (user_id)
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
