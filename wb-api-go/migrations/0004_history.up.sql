create table world.entity_history (
    id uuid not null default uuid_generate_v4(),
    entity_id uuid not null,
    historic_value jsonb not null,
    created_at timestamptz not null default now(),
    constraint pk_entityhistory_id primary key (id),
    constraint fk_entityhistory_entities foreign key (entity_id) references world.entities (id)
)