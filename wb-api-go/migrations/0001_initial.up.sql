create schema if not exists world;

-- create types
create table if not exists world.types
(
    id uuid not null default uuid_generate_v4(),
    parent_id uuid not null,
    wbtn text not null,
    type_name text not null,
    type_description text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_types_id primary key (id),
    constraint unq_types_wbtn unique (wbtn),
    constraint fk_types_types foreign key (parent_id) references world.types (id)
);

create index if not exists ix_types_parent_id on world.types using btree (parent_id asc nulls last);

-- create attribute_definitions
create table if not exists world.attribute_definitions
(
    id uuid not null default uuid_generate_v4(),
    wbatn text not null,
    attribute_name character varying(64)  not null,
    label character varying(64)  not null,
    data_type character varying(64)  not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_attributes_id primary key (id),
    constraint unq_attribute_definitions_wbatn unique (wbatn)
);

-- table: world.entities
create table if not exists world.entities
(
    id uuid not null default uuid_generate_v4(),
    type_id uuid not null,
    parent_id uuid not null,
    wbrn text  not null,
    entity_name text  not null,
    entity_description text  not null,
    notes text ,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entities_id primary key (id),
    constraint unq_entities_wbrn unique (wbrn),
    constraint fk_entities_entities foreign key (parent_id) references world.entities (id),
    constraint fk_entities_types foreign key (type_id) references world.types (id) 
);

 
create index if not exists ix_entities_parent_id on world.entities using btree (parent_id asc nulls last);
create index if not exists ix_entities_type_id on world.entities using btree (type_id asc nulls last);

-- table: world.entity_attributes
create table if not exists world.entity_attributes
(
    id uuid not null default uuid_generate_v4(),
    entity_id uuid not null,
    attribute_id uuid not null,
    attribute_value text  not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entityattributes_id primary key (id),
    constraint fk_entityattributes_attributes foreign key (attribute_id) references world.attribute_definitions (id),
    constraint fk_entityattributes_entities foreign key (entity_id) references world.entities (id) 
);
 
create index if not exists ix_entityattributes_attribute_id on world.entity_attributes using btree (attribute_id asc nulls last);
create index if not exists ix_entityattributes_entity_id on world.entity_attributes using btree (entity_id asc nulls last);

-- table: world.type_attributes
create table if not exists world.type_attributes
(
    type_id uuid not null,
    attribute_id uuid not null,
    ordinal integer not null,
    is_required boolean not null,
    constraint pk_type_attributes primary key (attribute_id, type_id),
    constraint fk_typeattributes_attributedefinitions foreign key (attribute_id) references world.attribute_definitions (id),
    constraint fk_typeattributes_types foreign key (type_id) references world.types (id) 
);
 
create index if not exists ix_typeattributes_attribute_id on world.type_attributes using btree (attribute_id asc nulls last);
create index if not exists ix_typeattributes_type_id on world.type_attributes using btree (type_id asc nulls last);