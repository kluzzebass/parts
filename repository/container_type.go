package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertContainerType(ctx context.Context, input model.NewContainerType) (*model.ContainerType, error) {
	sql := `
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
		RETURNING container_type_id, tenant_id, created_at, description
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
