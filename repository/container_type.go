package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertContainerType(ctx context.Context, input model.NewContainerType) (*model.ContainerType, error) {
	sql := `
		WITH
		inserted_rows AS (
			INSERT INTO
				container_type
			(
				container_type_id,
				tenant_id,
				description
			)
			VALUES
			(
				COALESCE($1, gen_random_uuid()),
				$2,
				$3
			)
			ON CONFLICT (container_type_id) DO UPDATE
			SET
				description = EXCLUDED.description
			WHERE
				container_type.description IS DISTINCT FROM EXCLUDED.description
			RETURNING
				container_type_id,
				tenant_id,
				created_at,
				description
		),
		selected_rows AS (
			SELECT
				container_type_id,
				tenant_id,
				created_at,
				description
			FROM
				container_type
			WHERE
				container_type_id = $1
		)
		SELECT
			COALESCE(ir.container_type_id, sr.container_type_id) AS container_type_id,
			COALESCE(ir.tenant_id, sr.tenant_id) AS tenant_id,
			COALESCE(ir.created_at, sr.created_at) AS created_at,
			COALESCE(ir.description, sr.description) AS description
		FROM
			inserted_rows ir
			FULL JOIN selected_rows sr USING (container_type_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.TenantID, input.Description)

	var dst *model.ContainerType
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListContainerTypes(ctx context.Context, ids *[]string) ([]*model.ContainerType, error) {
	sql := `
		WITH
		relevant_container_types AS (
			SELECT
				ct.container_type_id,
				ct.tenant_id,
				ct.created_at,
				ct.description
			FROM
				container_type ct
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND ct.container_type_id = ANY ($1::uuid[]))
		),
		relevant_containers AS (
			SELECT
				rct.container_type_id,
				array_agg(c.container_id) AS containers
			FROM
				relevant_container_types rct
				JOIN container c USING (container_type_id)
			GROUP BY
				rct.container_type_id
		)
		SELECT
			rct.container_type_id,
			rct.tenant_id,
			rct.created_at,
			rct.description,
			COALESCE(rc.containers, ARRAY[]::uuid[]) AS containers
		FROM
			relevant_container_types rct
			LEFT JOIN relevant_containers rc USING (container_type_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.ContainerType
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
