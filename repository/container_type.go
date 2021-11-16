package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
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

	var id string
	var tenantId string
	var createdAt time.Time
	var description string

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.TenantID, input.Description).Scan(&id, &tenantId, &createdAt, &description)

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
		rct.container_type_id AS id,
		rct.tenant_id,
		rct.created_at,
		rct.description,
		COALESCE(rc.containers, ARRAY[]::uuid[]) AS containers
	FROM
		relevant_container_types rct
		LEFT JOIN relevant_containers rc USING (container_type_id)
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
		var createdAt time.Time
		var description string
		var containers []string

		err := rows.Scan(&id, &tenantId, &createdAt, &description, &containers)
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
