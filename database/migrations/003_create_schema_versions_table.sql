create table schema_versions (
	id uuid primary key default uuid_generate_v4(),
	schema_id uuid not null references schemas(id),
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null
);

---- create above / drop below ----

drop table if exists schema_versions;
