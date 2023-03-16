-- Add down migration script here
-- Add up migration script here
alter table world.entities rename column community_id to community;
alter table world.communities rename column region_id to region;
alter table world.regions rename column country_id to country;