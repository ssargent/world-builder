-- Add down migration script here
drop function world.id;

drop index ix_entities_parent;
drop index ix_entities_type;
drop index ix_entities_wbrn;