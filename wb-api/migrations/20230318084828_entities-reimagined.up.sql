-- Add up migration script here

/*
create table if not exists world.entities (
    id uuid not null default (uuid_generate_v4()),
    wbrn text not null,
    entity_name text not null,
    entity_type text not null,
    community uuid null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entities_id primary key (id),
    constraint fk_entities_communties foreign key (community) references world.communities (id)
);
*/

drop table world.entities;

create table world.entities (
    id uuid not null default (uuid_generate_v4()),
    type_id uuid not null,
    parent_id uuid not null,
    wbrn text not null,
    entity_name text not null,
    entity_description text not null, 
    notes text null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entities_id primary key (id),
    constraint fk_entities_types foreign key (type_id) references world.types (id),
    constraint fk_entities_entities foreign key (parent_id) references world.entities (id)
);

insert into world.entities (id, type_id, parent_id, wbrn, entity_name, entity_description)
values (uuid_generate_v5(
            uuid_ns_url(),
            '/v1/worlds/ktos'
        ),
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/world'
        ),
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/worlds/ktos'
        ),
        'wb:ktos',
        'Ktos',
        'The planet Ktos'
);
 