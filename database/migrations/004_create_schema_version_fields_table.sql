create table schema_version_fields (
	id uuid primary key default uuid_generate_v4(),
	schema_version_id uuid not null references schema_versions(id),
	name varchar(100) not null,
	data_type varchar(50) not null,
	is_required boolean not null default false,
	default_value text null,
	validation_rules text[] null,
	ref_schema_version_id uuid null references schema_versions(id),
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null,
	check ( (data_type = 'schema') = (ref_schema_version_id is not null) )
);

---- create above / drop below ----

drop table if exists schema_version_fields;
