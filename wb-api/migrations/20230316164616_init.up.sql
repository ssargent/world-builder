-- Add up migration script here
create extension if not exists "uuid-ossp";

create table if not exists world.countries (
    id uuid not null default (uuid_generate_v4()),
    wbrn text not null,
    country_name text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_countries_id primary key (id) 
);

create table if not exists world.regions (
    id uuid not null default (uuid_generate_v4()),
    wbrn text not null,
    region_name text not null,
    country uuid null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_regions_id primary key (id),
    constraint fk_regions_countries foreign key (country) references world.countries (id)
);

create table if not exists world.communities (
    id uuid not null default (uuid_generate_v4()),
    wbrn text not null,
    community_name text not null,
    region uuid null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_communities_id primary key (id),
    constraint fk_communities_regions foreign key (region) references world.regions (id)
);

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