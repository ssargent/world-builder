-- Add up migration script here
create table if not exists world.types (
    id uuid not null default (uuid_generate_v4()),
    parent_id uuid not null,
    wbtn text not null,
    type_name text not null,
    type_description text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_types_id primary key (id)
);

insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/world'
        ),
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/world'
        ),
        'wbtn:world',
        'World',
        'A world'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/country'
        ),
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/world'
        ),
        'wbtn:country',
        'Country',
        'A country or kingdom'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/region'
        ),
         uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/country'
        ),
        'wbtn:region',
        'Region',
        'A region or province'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/community'
        ),
         uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/region'
        ),
        'wbtn:community',
        'Community',
        'A rural area or city'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/district'
        ),
         uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/community'
        ),
        'wbtn:district',
        'District',
        'A district within a municipality or rural area'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/location'
        ),
         uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/district'
        ),
        'wbtn:location',
        'Location',
        'A location with a specific address'
    );
insert into world.types (id, parent_id, wbtn, type_name, type_description)
values (
        uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/building'
        ),
         uuid_generate_v5(
            uuid_ns_url(),
            '/v1/types/district'
        ),
        'wbtn:building',
        'Building',
        'A specific building at a location'
    );

alter table world.types add constraint fk_types_types foreign key (parent_id) references world.types (id);

alter table world.entities drop column entity_type;
alter table world.entities add type_id uuid not null;
alter table world.entities add constraint fk_entities_type foreign key (type_id) references world.types (id);

