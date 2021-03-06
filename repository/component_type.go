package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertComponentType(ctx context.Context, input model.NewComponentType) (*model.ComponentType, error) {
	sql := `
		WITH
		inserted_rows AS (
			INSERT INTO
				component_type
			(
				component_type_id,
				tenant_id,
				description
			)
			VALUES
			(
				COALESCE($1, gen_random_uuid()),
				$2,
				$3
			)
			ON CONFLICT (component_type_id) DO UPDATE
			SET
				description = EXCLUDED.description
			WHERE
				component_type.description IS DISTINCT FROM EXCLUDED.description
			RETURNING
				component_type_id,
				tenant_id,
				created_at,
				description
		),
		selected_rows AS (
			SELECT
				component_type_id,
				tenant_id,
				created_at,
				description
			FROM
				component_type
			WHERE
				component_type_id = $1
		)
		SELECT
			COALESCE(ir.component_type_id, sr.component_type_id) AS component_type_id,
			COALESCE(ir.tenant_id, sr.tenant_id) AS tenant_id,
			COALESCE(ir.created_at, sr.created_at) AS created_at,
			COALESCE(ir.description, sr.description) AS description
		FROM
			inserted_rows ir
			FULL JOIN selected_rows sr USING (component_type_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.TenantID, input.Description)

	var dst *model.ComponentType
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListComponentTypes(ctx context.Context, ids *[]string) ([]*model.ComponentType, error) {
	sql := `
		WITH
		relevant_component_types AS (
			SELECT
				ct.component_type_id,
				ct.tenant_id,
				ct.created_at,
				ct.description
			FROM
				component_type ct
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND ct.component_type_id = ANY ($1::uuid[]))
		),
		relevant_components AS (
			SELECT
				rct.component_type_id,
				array_agg(c.component_id) AS components
			FROM
				relevant_component_types rct
				JOIN component c USING (component_type_id)
			GROUP BY
				rct.component_type_id
		)
		SELECT
			rct.component_type_id,
			rct.tenant_id,
			rct.created_at,
			rct.description,
			COALESCE(rc.components, ARRAY[]::uuid[]) AS components
		FROM
			relevant_component_types rct
			LEFT JOIN relevant_components rc USING (component_type_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.ComponentType
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
