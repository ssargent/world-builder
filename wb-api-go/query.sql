-- name: GetTypeByID :one
select * from world.types
where id = $1;

-- name: GetTypeByWBTN :one
select * from world.types
where wbtn = $1;

-- name: GetEntityAttributes :many
select * from 
world.entity_attributes 
where entity_id = $1;

-- name: GetAttributesForType :many
select ad.* from
world.attribute_definitions ad inner join world.type_attributes ta
on ad.id = ta.attribute_id 
inner join world.types t on t.id = ta.type_id
where ta.type_id = $1;

-- name: CreateType :one
insert into world.types
(parent_id, wbtn, type_name, type_description)
values 
($1, $2, $3, $4)
returning *;

-- name: CreateAttributeDefinition :one
insert into world.attribute_definitions
(wbatn, attribute_name, label, data_type)
values
($1, $2, $3, $4)
returning *;

-- name: GetEntityByWBRN :one
select e.* from world.entities e
where e.wbrn = $1;

-- name: GetEntitiesByWBRN :many
select e.* from world.entities e
where e.wbrn like $1;

-- name: CreateEntity :one
insert into world.entities
(id, type_id, parent_id, wbrn, entity_name, entity_description, notes)
values 
($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetEntityAssociationsForEntity :many
select id, entity_one, entity_two, type_id, effective_start_date, effective_end_date
from world.entity_associations
where (entity_one = $1 or entity_two = $1);

-- name: CreateEntityAssociation :one
insert into world.entity_associations
(entity_one, entity_two, type_id, effective_start_date, effective_end_date)
values 
($1, $2, $3, $4, $5) 
returning *;