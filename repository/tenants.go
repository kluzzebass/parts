package repository

import (
	"context"
	"parts/graph/model"

	"github.com/randallmlough/pgxscan"
)

func (repo *Repository) UpsertTenant(ctx context.Context, input model.NewTenant) (*model.Tenant, error) {
	sql := `
		INSERT INTO
			tenant
		(
			tenant_id,
			name
		)
		VALUES
		(
			COALESCE($1, gen_random_uuid()),
			$2
		)
		ON CONFLICT (tenant_id) DO UPDATE
		SET
			name = EXCLUDED.name
		RETURNING tenant_id, created_at, name
	`
	rows, _ := repo.pool.Query(ctx, sql, input.ID, input.Name)

	var dst *model.Tenant
	if err := pgxscan.NewScanner(rows).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo *Repository) ListTenants(ctx context.Context, ids *[]string) ([]*model.Tenant, error) {
	sql := `
		WITH
		relevant_tenants AS (
			SELECT
			t.tenant_id,
			t.created_at,
			t.name
		FROM
			tenant t
		WHERE
			$1::uuid[] IS NULL
			OR ($1::uuid[] IS NOT NULL AND t.tenant_id = ANY ($1::uuid[]))
		),
		relevant_users AS (
			SELECT
				t.tenant_id,
				array_agg(u.user_id) as users
			FROM
				tenant t
				JOIN "user" u USING (tenant_id)
			GROUP BY
				t.tenant_id
		),
		relevant_container_types AS (
			SELECT
				t.tenant_id,
				array_agg(ct.container_type_id) as container_types
			FROM
				tenant t
				JOIN container_type ct USING (tenant_id)
			GROUP BY
				t.tenant_id
		),
		relevant_component_types AS (
			SELECT
				t.tenant_id,
				array_agg(ct.component_type_id) as component_types
			FROM
				tenant t
				JOIN component_type ct USING (tenant_id)
			GROUP BY
				t.tenant_id
		)
		SELECT
			rt.tenant_id,
			rt.created_at,
			rt.name,
			COALESCE(ru.users, ARRAY[]::uuid[]) AS users,
			COALESCE(rcnt.container_types, ARRAY[]::uuid[]) as container_types,
			COALESCE(rcmt.component_types, ARRAY[]::uuid[]) as component_types
		FROM
			relevant_tenants rt
			LEFT JOIN relevant_users ru USING (tenant_id)
			LEFT JOIN relevant_container_types rcnt USING (tenant_id)
			LEFT JOIN relevant_component_types rcmt USING (tenant_id)
	`

	rows, _ := repo.pool.Query(ctx, sql, ids)

	var dst []*model.Tenant
	if err := pgxscan.NewScanner(rows, pgxscan.ErrNoRowsQuery(false)).Scan(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
