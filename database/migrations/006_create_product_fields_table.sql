create table product_fields (
	id uuid primary key default uuid_generate_v4(),
	product_id uuid not null references products(id),
	schema_version_field_id uuid not null references schema_version_fields(id),
	value text null,
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null
);

---- create above / drop below ----

drop table if exists product_fields;
