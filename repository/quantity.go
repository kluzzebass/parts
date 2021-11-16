package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
	"time"
)

func (repo *Repository) UpsertQuantity(ctx context.Context, input model.NewQuantity) (*model.Quantity, error) {
	sql := `
		INSERT INTO
			quantity
		(
			quantity_id,
			container_id,
			component_id,
			quantity
		)
		VALUES
		(
			COALESCE($1, gen_random_uuid()),
			$2,
			$3,
			$4
		)
		ON CONFLICT (quantity_id) DO UPDATE
		SET
			quantity = EXCLUDED.quantity
		RETURNING quantity_id, container_id, component_id, created_at, quantity
	`

	var id string
	var containerId string
	var componentId string
	var createdAt time.Time
	var quantity int

	err := repo.pool.QueryRow(ctx, sql, input.ID, input.ContainerID, input.ComponentID, input.Quantity).Scan(&id, &containerId, &componentId, &createdAt, &quantity)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.Quantity{
		ID:          id,
		ContainerID: containerId,
		ComponentID: componentId,
		CreatedAt:   createdAt,
		Quantity:    quantity,
	}, err
}

func (repo *Repository) ListQuantities(ctx context.Context, ids *[]string) ([]*model.Quantity, error) {
	sql := `
		WITH
		relevant_quantities AS (
			SELECT
				q.quantity_id,
				q.container_id,
				q.component_id,
				q.created_at,
				q.quantity
			FROM
				quantity q
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND q.quantity_id = ANY ($1::uuid[]))
		)
		SELECT
			rq.quantity_id AS id,
			rq.container_id,
			rq.component_id,
			rq.created_at,
			rq.quantity
		FROM
			relevant_quantities rq
	`

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var quantities []*model.Quantity

	for rows.Next() {
		var id string
		var containerId string
		var componentId string
		var createdAt time.Time
		var quantity int

		err := rows.Scan(&id, &containerId, &componentId, &createdAt, &quantity)
		if err != nil {
			return nil, err
		}

		quantities = append(quantities, &model.Quantity{
			ID:          id,
			ContainerID: containerId,
			ComponentID: componentId,
			CreatedAt:   createdAt,
			Quantity:    quantity,
		})
	}

	return quantities, err
}
