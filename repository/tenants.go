package repository

import (
	"context"
	"fmt"
	"os"
	"parts/graph/model"
)

func (repo *Repository) CreateTenant(ctx context.Context, nt model.NewTenant) (*model.Tenant, error) {
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
		RETURNING tenant_id, created_at::text, name
	`
	var id string
	var createdAt string
	var name string

	err := repo.pool.QueryRow(ctx, sql, nt.ID, nt.Name).Scan(&id, &createdAt, &name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &model.Tenant{
		ID:        id,
		CreatedAt: createdAt,
		Name:      name,
	}, err
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
			rt.tenant_id AS id,
			rt.created_at::text,
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

	rows, err := repo.pool.Query(ctx, sql, ids)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var tenants []*model.Tenant

	for rows.Next() {
		var id string
		var createdAt string
		var name string
		var users []string
		var containerTypes []string
		var componentTypes []string

		err := rows.Scan(&id, &createdAt, &name, &users, &containerTypes, &componentTypes)

		if err != nil {
			return nil, err
		}

		tenants = append(tenants, &model.Tenant{
			ID:               id,
			CreatedAt:        createdAt,
			Name:             name,
			UserIDs:          users,
			ContainerTypeIDs: containerTypes,
			ComponentTypeIDs: componentTypes,
		})
	}

	return tenants, err
}
