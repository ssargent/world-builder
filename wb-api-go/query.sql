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