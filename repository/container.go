package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertContainer(ctx context.Context, input model.NewContainer) (*model.Container, error) {
	sql := `
	WITH
	inserted_rows AS (
			INSERT INTO
				container
			(
				container_id,
				tenant_id,
				parent_container_id,
				container_type_id,
				description
			)
			VALUES
			(
				COALESCE($1, gen_random_uuid()),
				$2,
				$3,
				$4,
				$5
			)
			ON CONFLICT (container_id) DO UPDATE
			SET
				parent_container_id = EXCLUDED.parent_container_id,
				container_type_id = EXCLUDED.container_type_id,
				description = EXCLUDED.description
			WHERE
				container.parent_container_id IS DISTINCT FROM EXCLUDED.parent_container_id
				OR container.container_type_id IS DISTINCT FROM EXCLUDED.container_type_id
				OR container.description IS DISTINCT FROM EXCLUDED.description
			RETURNING
				container_id,
				tenant_id,
				parent_container_id,
				container_type_id,
				created_at,
				description
		),
		selected_rows AS (
			SELECT
				container_id,
				tenant_id,
				parent_container_id,
				container_type_id,
				created_at,
				description
			FROM
				container
			WHERE
				container_id = $1
		)
		SELECT
			COALESCE(ir.container_id, sr.container_id) AS container_id,
			COALESCE(ir.tenant_id, sr.tenant_id) AS tenant_id,
			COALESCE(ir.parent_container_id, sr.parent_container_id) AS parent_container_id,
			COALESCE(ir.container_type_id, sr.container_type_id) AS container_type_id,
			COALESCE(ir.created_at, sr.created_at) AS created_at,
			COALESCE(ir.description, sr.description) AS description
		FROM
			inserted_rows ir
			FULL JOIN selected_rows sr USING (container_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.TenantID, input.ParentID, input.ContainerTypeID, input.Description)

	var dst *model.Container
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListContainers(ctx context.Context, ids *[]string) ([]*model.Container, error) {
	sql := `
		WITH
		relevant_containers AS (
			SELECT
				c.container_id,
				c.tenant_id,
				c.parent_container_id,
				c.container_type_id,
				c.created_at,
				c.description
			FROM
				container c
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND c.container_id = ANY ($1::uuid[]))
		),
		relevant_children AS (
			SELECT
				rc.container_id,
				array_agg(c.container_id) AS children
			FROM
				relevant_containers rc
				JOIN container c ON rc.container_id = c.parent_container_id
			GROUP BY
				rc.container_id
		),
		relevant_quantities AS (
			SELECT
				rc.container_id,
				array_agg(q.quantity_id) AS quantities
			FROM
				relevant_containers rc
				JOIN quantity q USING (container_id)
			GROUP BY
				rc.container_id
		)
		SELECT
			rc.container_id,
			rc.tenant_id,
			rc.parent_container_id,
			rc.container_type_id,
			rc.created_at,
			rc.description,
			COALESCE(rch.children, ARRAY[]::uuid[]) AS children,
			COALESCE(rq.quantities, ARRAY[]::uuid[]) AS quantities
		FROM
			relevant_containers rc
			LEFT JOIN relevant_children rch USING (container_id)
			LEFT JOIN relevant_quantities rq USING (container_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.Container
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
