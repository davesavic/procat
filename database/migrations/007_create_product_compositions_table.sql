create table product_compositions (
	id uuid primary key default uuid_generate_v4(),
	parent_product_id uuid references products(id),
	schema_version_field_id uuid references schema_version_fields(id),
	child_product_id uuid references products(id),
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null
);

---- create above / drop below ----

drop table if exists product_compositions;
