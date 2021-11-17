package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertQuantity(ctx context.Context, input model.NewQuantity) (*model.Quantity, error) {
	sql := `
		WITH
		inserted_rows AS (
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
			WHERE
				quantity.quantity IS DISTINCT FROM EXCLUDED.quantity
			RETURNING
				quantity_id,
				container_id,
				component_id,
				created_at,
				quantity
		),
		selected_rows AS (
			SELECT
				quantity_id,
				container_id,
				component_id,
				created_at,
				quantity
			FROM
				quantity
			WHERE
			quantity_id = $1
		),
		inserted_quantity_changes AS (
			INSERT INTO
				quantity_change
			(
				quantity_id,
				quantity
			)
			SELECT
				quantity_id,
				quantity
			FROM
				inserted_rows
			RETURNING
				quantity_change_id
		)
		SELECT
			COALESCE(ir.quantity_id, sr.quantity_id) AS quantity_id,
			COALESCE(ir.container_id, sr.container_id) AS container_id,
			COALESCE(ir.component_id, sr.component_id) AS component_id,
			COALESCE(ir.created_at, sr.created_at) AS created_at,
			COALESCE(ir.quantity, sr.quantity) AS quantity
		FROM
			inserted_rows ir
			FULL JOIN selected_rows sr USING (quantity_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.ContainerID, input.ComponentID, input.Quantity)

	var dst *model.Quantity
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
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
			rq.quantity_id,
			rq.container_id,
			rq.component_id,
			rq.created_at,
			rq.quantity
		FROM
			relevant_quantities rq
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.Quantity
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
