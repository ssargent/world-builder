-- Add up migration script here
create or replace function world.id(typestring text) returns uuid as $$
	begin
		return uuid_generate_v5(uuid_ns_url(), typestring);
	end; 
$$ language plpgsql;

create index ix_entities_type on world.entities (type_id);
create index ix_entities_wbrn on world.entities (wbrn);
create index ix_entities_parent on world.entities (parent_id);