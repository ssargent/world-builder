-- Add up migration script here
create table world.entity_attributes (
    id uuid not null default (uuid_generate_v4()),
    entity_id uuid not null,
    attribute_name varchar(64) not null,
    label varchar(64) not null,
    data_type varchar(64) not null,
    attribute_value text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_attributes_id primary key (id),
    constraint fk_attributes_entities foreign key (entity_id) references world.entities (id)
)