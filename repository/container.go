package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
)

func (repo *Repository) UpsertContainer(ctx context.Context, input model.NewContainer) (*model.Container, error) {
	sql := `
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
		RETURNING container_id, tenant_id, parent_container_id, container_type_id, created_at, description
	`

	var id string
	var tenantId string
	var parentId *string
	var containerTypeId string
	var createdAt time.Time
	var description string

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.TenantID, input.ParentID, input.ContainerTypeID, input.Description).Scan(&id, &tenantId, &parentId, &containerTypeId, &createdAt, &description)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.Container{
		ID:              id,
		TenantID:        tenantId,
		ParentID:        parentId,
		ContainerTypeID: containerTypeId,
		CreatedAt:       createdAt,
		Description:     description,
	}, err
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
		rc.container_id AS id,
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

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var containers []*model.Container

	for rows.Next() {
		var id string
		var tenantId string
		var parentId *string
		var containerTypeId string
		var createdAt time.Time
		var description string
		var children []string
		var quantities []string

		err := rows.Scan(&id, &tenantId, &parentId, &containerTypeId, &createdAt, &description, &children, &quantities)
		if err != nil {
			return nil, err
		}

		containers = append(containers, &model.Container{
			ID:              id,
			TenantID:        tenantId,
			ParentID:        parentId,
			ContainerTypeID: containerTypeId,
			CreatedAt:       createdAt,
			Description:     description,
			ChildIDs:        children,
			QuantityIDs:     quantities,
		})
	}

	return containers, err
}
