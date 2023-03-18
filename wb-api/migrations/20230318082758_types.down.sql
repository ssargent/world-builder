-- Add down migration script here
alter table world.entities drop constraint fk_entities_type;
alter table world.types drop constraint fk_types_types;

alter table world.entities drop column type_id;
alter table world.entitites add entity_type text not null default('unknown');

drop table world.types;