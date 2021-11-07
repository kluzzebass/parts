

CREATE TABLE tenant (
  tenant_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at timestamptz NOT NULL DEFAULT now(),
  name text NOT NULL,
  UNIQUE (name)
);

CREATE TABLE "user" (
  user_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id uuid NOT NULL REFERENCES tenant (tenant_id),
  created_at timestamptz NOT NULL DEFAULT now(),
  name text NOT NULL,
  email text NOT NULL,
  UNIQUE (tenant_id, email)
);

CREATE INDEX user_tenant_id_idx ON "user" (tenant_id);

CREATE TABLE container_type (
  container_type_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id uuid NOT NULL REFERENCES tenant (tenant_id),
  created_at timestamptz NOT NULL DEFAULT now(),
  description text NOT NULL,
  UNIQUE (tenant_id, description)
);

CREATE INDEX container_type_tenant_id_idx ON container_type (tenant_id);

CREATE TABLE container (
  container_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id uuid NOT NULL REFERENCES tenant (tenant_id),
  parent_container_id uuid REFERENCES container (container_id),
  container_type_id uuid NOT NULL REFERENCES container_type (container_type_id),
  created_at timestamptz NOT NULL DEFAULT now(),
  description text NOT NULL,
  UNIQUE (tenant_id, description)
);

CREATE INDEX container_tenant_id_idx ON container (tenant_id);
CREATE INDEX container_parent_container_id_idx ON container (parent_container_id);
CREATE INDEX container_container_type_id_idx ON container_type (container_type_id);


CREATE TABLE component_type (
  component_type_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id uuid NOT NULL REFERENCES tenant (tenant_id),
  created_at timestamptz NOT NULL DEFAULT now(),
  description text NOT NULL,
  UNIQUE (tenant_id, description)
);

CREATE INDEX component_type_tenant_id_idx ON container (tenant_id);


CREATE TABLE component (
  component_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id uuid NOT NULL REFERENCES tenant (tenant_id),
  component_type_id uuid NOT NULL REFERENCES component_type (component_type_id),
  created_at timestamptz NOT NULL DEFAULT now(),
  description text NOT NULL,
  UNIQUE (tenant_id, description)
);

CREATE INDEX component_tenant_id_idx ON component (tenant_id);
CREATE INDEX component_component_type_id_idx ON component (component_type_id);

CREATE TABLE container_component (
  container_id uuid NOT NULL REFERENCES container (container_id),
  component_id uuid NOT NULL REFERENCES component (component_id),
  quantity int NOT NULL default 0 CHECK (quantity >= 0),
  PRIMARY KEY (container_id, component_id)
);

CREATE INDEX container_component_container_id ON container_component (container_id);
CREATE INDEX container_component_component_id ON container_component (component_id);




