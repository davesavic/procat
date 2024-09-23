create table schemas (
	id uuid primary key default uuid_generate_v4(),
	name text not null,
	description text null,
	created_at timestamp with time zone default timezone('utc', now()),
	updated_at timestamp with time zone default timezone('utc', now()),
	deleted_at timestamp with time zone null
);

---- create above / drop below ----

drop table if exists schemas;
