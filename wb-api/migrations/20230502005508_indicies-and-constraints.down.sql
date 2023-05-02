-- Add down migration script here
alter table world.attribute_definitions drop constraint unq_attribute_definitions_wbatn;
drop index  uix_attribute_definitions_wbatn;

alter table world.entities drop constraint unq_entities_wbrn;
drop index  uix_entities_wbrn;

alter table world.entities drop constraint unq_types_wbtn;
drop index  uix_types_wbtn;

drop index  ix_types_parent_id;
drop index  ix_entities_parent_id;
drop index  ix_entities_type_id;
drop index  ix_typeattributes_type_id;
drop index  ix_typeattributes_attribute_id;
drop index  ix_entityattributes_entity_id;
drop index  ix_entityattributes_attribute_id;