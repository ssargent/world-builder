-- Add up migration script here
 
drop table world.entity_attributes;

create table world.attribute_definitions (
    id uuid not null default (uuid_generate_v4()),
    wbatn text not null,
    attribute_name varchar(64) not null,
    label varchar(64) not null,
    data_type varchar(64) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_attributes_id primary key (id),
    constraint unq_attribute_definitions_wbatn unique(wbatn)
); 

create table world.type_attributes (
    type_id uuid not null,
    attribute_id uuid not null,
    ordinal int not null,
    is_required boolean not null,
    constraint pk_type_attributes primary key (attribute_id, type_id),
    constraint fk_typeattributes_types foreign key (type_id) references world.types (id),
    constraint fk_typeattributes_attributedefinitions foreign key (attribute_id) references world.attribute_definitions (id)
);

create index ix_type_attributes_type_id on world.type_attributes (type_id);
create index ix_type_attributes_attribute_id on world.type_attributes (attribute_id);

create table world.entity_attributes (
    id uuid not null default (uuid_generate_v4()),
    entity_id uuid not null,
    attribute_id uuid not null,
    attribute_value text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entityattributes_id primary key (id),
    constraint fk_entityattributes_entities foreign key (entity_id) references world.entities (id),
    constraint fk_entityattributes_attributes foreign key (attribute_id) references world.attribute_definitions (id)
);

create index ix_entityattributes_entity_id on world.entity_attributes (entity_id);
create index ix_entityattributes_attribute_id on world.entity_attributes (attribute_id);