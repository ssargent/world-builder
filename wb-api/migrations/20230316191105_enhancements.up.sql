-- Add up migration script here
alter table world.entities rename column community to community_id;
alter table world.communities rename column region to region_id;
alter table world.regions rename column country to country_id;

