package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
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

	var id string
	var tenantId string
	var componentTypeId string
	var createdAt time.Time
	var description string

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.TenantID, input.ComponentTypeID, input.Description).Scan(&id, &tenantId, &componentTypeId, &createdAt, &description)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.Component{
		ID:              id,
		TenantID:        tenantId,
		ComponentTypeID: componentTypeId,
		CreatedAt:       createdAt,
		Description:     description,
	}, err
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
			rc.component_type_id AS id,
			rc.created_at,
			rc.description,
			COALESCE(rq.quantities, ARRAY[]::uuid[]) AS quantities
		FROM
			relevant_components rc
			LEFT JOIN relevant_quantities rq USING (component_id)
	`

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var components []*model.Component

	for rows.Next() {
		var id string
		var tenantId string
		var componentTypeId string
		var createdAt time.Time
		var description string
		var quantities []string

		err := rows.Scan(&id, &tenantId, &componentTypeId, &createdAt, &description, &quantities)
		if err != nil {
			return nil, err
		}

		components = append(components, &model.Component{
			ID:              id,
			TenantID:        tenantId,
			ComponentTypeID: componentTypeId,
			CreatedAt:       createdAt,
			Description:     description,
			QuantityIDs:     quantities,
		})
	}

	return components, err
}
