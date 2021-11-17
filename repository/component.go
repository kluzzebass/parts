package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertComponent(ctx context.Context, input model.NewComponent) (*model.Component, error) {
	sql := `
		INSERT INTO
			component
		(
			component_id,
			tenant_id,
			component_type_id,
			description
		)
		VALUES
		(
			COALESCE($1, gen_random_uuid()),
			$2,
			$3,
			$4
		)
		ON CONFLICT (component_id) DO UPDATE
		SET
			component_type_id = EXCLUDED.component_type_id,
			description = EXCLUDED.description
		RETURNING component_id, tenant_id, component_type_id, created_at, description
	`

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.TenantID, input.ComponentTypeID, input.Description)

	var dst *model.Component
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListComponents(ctx context.Context, ids *[]string) ([]*model.Component, error) {
	sql := `
		WITH
		relevant_components AS (
			SELECT
				c.component_id,
				c.tenant_id,
				c.component_type_id,
				c.created_at,
				c.description
			FROM
				component c
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND c.component_id = ANY ($1::uuid[]))
		),
		relevant_quantities AS (
			SELECT
				rc.component_id,
				array_agg(q.quantity_id) AS quantities
			FROM
				relevant_components rc
				JOIN quantity q USING (component_id)
			GROUP BY
				rc.component_id
		)
		SELECT
			rc.component_id,
			rc.tenant_id,
			rc.component_type_id,
			rc.created_at,
			rc.description,
			COALESCE(rq.quantities, ARRAY[]::uuid[]) AS quantities
		FROM
			relevant_components rc
			LEFT JOIN relevant_quantities rq USING (component_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.Component
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
