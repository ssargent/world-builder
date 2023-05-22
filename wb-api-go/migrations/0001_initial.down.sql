create schema if not exists world;

drop index if exists world.ix_typeattributes_type_id;
drop index if exists world.ix_typeattributes_attribute_id; 
drop table if exists world.type_attributes;

drop index if exists world.ix_entityattributes_entity_id;
drop index if exists world.ix_entityattributes_attribute_id;
drop table if exists world.entity_attributes;

drop index if exists world.ix_entities_parent_id;
drop index if exists world.ix_entities_type_id;
drop table if exists world.entities;

drop table if exists world.attribute_definitions;

drop index if exists world.ix_types_parent_id;
drop table if exists world.types;