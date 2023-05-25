
CREATE SCHEMA IF NOT EXISTS world
    AUTHORIZATION wb;
    
CREATE TABLE IF NOT EXISTS world.attribute_definitions
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    wbatn text COLLATE pg_catalog."default" NOT NULL,
    attribute_name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    label character varying(64) COLLATE pg_catalog."default" NOT NULL,
    data_type character varying(64) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT pk_attributes_id PRIMARY KEY (id),
    CONSTRAINT unq_attribute_definitions_wbatn UNIQUE (wbatn)
);

CREATE TABLE IF NOT EXISTS world.types
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    parent_id uuid NOT NULL,
    wbtn text COLLATE pg_catalog."default" NOT NULL,
    type_name text COLLATE pg_catalog."default" NOT NULL,
    type_description text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT pk_types_id PRIMARY KEY (id),
    CONSTRAINT unq_types_wbtn UNIQUE (wbtn),
    CONSTRAINT fk_types_types FOREIGN KEY (parent_id)
        REFERENCES world.types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS world.entities
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    type_id uuid NOT NULL,
    parent_id uuid NOT NULL,
    wbrn text COLLATE pg_catalog."default" NOT NULL,
    entity_name text COLLATE pg_catalog."default" NOT NULL,
    entity_description text COLLATE pg_catalog."default" NOT NULL,
    notes text COLLATE pg_catalog."default",
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT pk_entities_id PRIMARY KEY (id),
    CONSTRAINT unq_entities_wbrn UNIQUE (wbrn),
    CONSTRAINT fk_entities_entities FOREIGN KEY (parent_id)
        REFERENCES world.entities (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_entities_types FOREIGN KEY (type_id)
        REFERENCES world.types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS world.entity_attributes
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    entity_id uuid NOT NULL,
    attribute_id uuid NOT NULL,
    attribute_value text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT pk_entityattributes_id PRIMARY KEY (id),
    CONSTRAINT fk_entityattributes_attributes FOREIGN KEY (attribute_id)
        REFERENCES world.attribute_definitions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_entityattributes_entities FOREIGN KEY (entity_id)
        REFERENCES world.entities (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS world.type_attributes
(
    type_id uuid NOT NULL,
    attribute_id uuid NOT NULL,
    ordinal integer NOT NULL,
    is_required boolean NOT NULL,
    CONSTRAINT pk_type_attributes PRIMARY KEY (attribute_id, type_id),
    CONSTRAINT fk_typeattributes_attributedefinitions FOREIGN KEY (attribute_id)
        REFERENCES world.attribute_definitions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_typeattributes_types FOREIGN KEY (type_id)
        REFERENCES world.types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

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

create table world.entity_history (
    id uuid not null default uuid_generate_v4(),
    entity_id uuid not null,
    historic_value jsonb not null,
    created_at timestamptz not null default now(),
    constraint pk_entityhistory_id primary key (id),
    constraint fk_entityhistory_entities foreign key (entity_id) references world.entities (id)
);