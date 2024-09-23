create table products (
	id uuid primary key default uuid_generate_v4(),
	schema_version_id uuid not null references schema_versions(id),
	name text not null,
	description text null,
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null
);

---- create above / drop below ----

drop table if exists products;
