package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
)

func (repo *Repository) CreateComponentType(ctx context.Context, nt model.NewComponentType) (*model.ComponentType, error) {
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
		RETURNING component_type_id, tenant_id, created_at::text, description
	`

	var id string
	var tenantId string
	var createdAt string
	var description string

	err := repo.pool.QueryRow(ctx, sql, nt.ID, nt.TenantID, nt.Description).Scan(&id, &tenantId, &createdAt, &description)

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
		SELECT
			component_type_id AS id,
			tenant_id,
			created_at::text,
			description
		FROM
			component_type
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND component_type_id = ANY ($1::uuid[]))
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
		var createdAt string
		var description string
		err := rows.Scan(&id, &tenantId, &createdAt, &description)
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
