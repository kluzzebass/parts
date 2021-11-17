package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) ListQuantityChanges(ctx context.Context, ids *[]string) ([]*model.QuantityChange, error) {
	sql := `
		WITH
		relevant_quantity_changes AS (
			SELECT
				qc.quantity_change_id,
				qc.quantity_id,
				qc.created_at,
				qc.amount
			FROM
				quantity_change qc
			WHERE
				$1::uuid[] IS NULL
				OR ($1::uuid[] IS NOT NULL AND qc.quantity_change_id = ANY ($1::uuid[]))
		)
		SELECT
			rqc.quantity_change_id,
			rqc.quantity_id,
			rqc.created_at,
			rqc.amount
		FROM
			relevant_quantity_changes rqc
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.QuantityChange
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
