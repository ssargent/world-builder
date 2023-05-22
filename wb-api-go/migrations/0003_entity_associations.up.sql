create table if not exists world.entity_associations (
    id uuid not null default uuid_generate_v4(),
    entity_one uuid not null,
    entity_two uuid not null,
    type_id uuid not null,
    effective_start_date timestamp with time zone default now(),
    effective_end_date timestamp with time zone,
    constraint pk_entityassociations_id primary key (id),
    constraint fk_entityassociations_entities_one foreign key (entity_one) references world.entities (id),
    constraint fk_entityassociations_entities_two foreign key (entity_two) references world.entities (id),
    constraint fk_entityassociations_types foreign key (type_id) references world.types (id)
);

create index if not exists ix_entityassociations_association_type on world.entity_associations using btree (type_id asc nulls last);