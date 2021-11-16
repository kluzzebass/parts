package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
)

func (repo *Repository) CreateContainerType(ctx context.Context, nt model.NewContainerType) (*model.ContainerType, error) {
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
		RETURNING container_type_id, tenant_id, created_at::text, description
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

	return &model.ContainerType{
		ID:          id,
		TenantID:    tenantId,
		CreatedAt:   createdAt,
		Description: description,
	}, err
}

func (repo *Repository) ListContainerTypes(ctx context.Context, ids *[]string) ([]*model.ContainerType, error) {
	sql := `
		SELECT
			container_type_id AS id,
			tenant_id,
			created_at::text,
			description
		FROM
			container_type
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND container_type_id = ANY ($1::uuid[]))
	`

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var containerTypes []*model.ContainerType

	for rows.Next() {
		var id string
		var tenantId string
		var createdAt string
		var description string
		err := rows.Scan(&id, &tenantId, &createdAt, &description)
		if err != nil {
			return nil, err
		}

		containerTypes = append(containerTypes, &model.ContainerType{
			ID:          id,
			TenantID:    tenantId,
			CreatedAt:   createdAt,
			Description: description,
		})
	}

	return containerTypes, err
}
