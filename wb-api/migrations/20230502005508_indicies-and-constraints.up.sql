-- Add up migration script here

create index if not exists ix_types_parent_id on world.types (parent_id);
create index if not exists   ix_entities_parent_id on world.entities (parent_id);
create index if not exists   ix_entities_type_id on world.entities (type_id);
create index if not exists   ix_typeattributes_type_id on world.type_attributes (type_id);
create index if not exists   ix_typeattributes_attribute_id on world.type_attributes (attribute_id);
create index if not exists   ix_entityattributes_entity_id on world.entity_attributes (entity_id);
create index if not exists   ix_entityattributes_attribute_id on world.entity_attributes (attribute_id);
 
create unique index if not exists   uix_types_wbtn on world.types (wbtn);
alter table world.types add constraint unq_types_wbtn unique using index uix_types_wbtn;

create unique index if not exists   uix_entities_wbrn on world.entities (wbrn);
alter table world.entities add constraint unq_entities_wbrn unique using index uix_entities_wbrn;

create unique index if not exists   uix_attribute_definitions_wbatn on world.attribute_definitions (wbatn);
alter table world.attribute_definitions add constraint unq_attribute_definitions_wbatn unique using index uix_attribute_definitions_wbatn;