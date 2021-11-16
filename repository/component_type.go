package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
)

func (repo *Repository) UpsertComponentType(ctx context.Context, input model.NewComponentType) (*model.ComponentType, error) {
	sql := `
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
		RETURNING component_type_id, tenant_id, created_at, description
	`

	var id string
	var tenantId string
	var createdAt time.Time
	var description string

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.TenantID, input.Description).Scan(&id, &tenantId, &createdAt, &description)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.ComponentType{
		ID:          id,
		TenantID:    tenantId,
		CreatedAt:   createdAt,
		Description: description,
	}, err
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
			rct.component_type_id AS id,
			rct.tenant_id,
			rct.created_at,
			rct.description,
			COALESCE(rc.components, ARRAY[]::uuid[]) AS components
		FROM
			relevant_component_types rct
			LEFT JOIN relevant_components rc USING (component_type_id)
	`

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var componentTypes []*model.ComponentType

	for rows.Next() {
		var id string
		var tenantId string
		var createdAt time.Time
		var description string
		var components []string

		err := rows.Scan(&id, &tenantId, &createdAt, &description, &components)
		if err != nil {
			return nil, err
		}

		componentTypes = append(componentTypes, &model.ComponentType{
			ID:          id,
			TenantID:    tenantId,
			CreatedAt:   createdAt,
			Description: description,
		})
	}

	return componentTypes, err
}
