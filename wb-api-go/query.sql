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

-- name: GetAttributeByWBATN :one
select * from world.attribute_definitions
where wbatn = $1;

-- name: CreateType :one
insert into world.types
(parent_id, wbtn, type_name, type_description)
values
($1, $2, $3, $4)
returning *;

-- name: CreateTypeAttribute :one
insert into world.type_attributes
(type_id, attribute_id, ordinal, is_required)
values ($1, $2, $3, $4)
returning *;

-- name: CreateAttributeDefinition :one
insert into world.attribute_definitions
(wbatn, attribute_name, label, data_type)
values
($1, $2, $3, $4)
returning *;

-- name: GetFullTypeAttributes :many
select ad.*, ta.ordinal, ta.is_required from world.attribute_definitions ad
inner join world.type_attributes ta
on ta.attribute_id = ad.id
where ta.type_id = $1;

-- name: GetEntity :one
select e.* from world.entities e where id = $1;


-- name: GetEntityByWBRN :one
select e.* from world.entities e
where e.wbrn = $1;

-- name: GetEntitiesByWBRN :many
select e.* from world.entities e
where e.wbrn like $1;

-- name: GetEntitiesByParent :many
select e.* from world.entities e
where e.parent_id = $1;

-- name: GetEntitiesByCriteria :many
select e.*
from world.entities e
	inner join world.types w on e.type_id = w.id
	inner join world.entities p on e.parent_id = p.id
where (w.wbtn = $1 or $1 = '')
and (p.wbrn = $2 or $2 = '');

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

-- name: CreateEntityHistory :one
insert into world.entity_history
(entity_id, historic_value)
values
($1, $2)
returning *;

-- name: GetEntityHistory :many
select id, entity_id, historic_value, created_at
from world.entity_history
where entity_id = $1
order by created_at;

-- name: GetEntityChildReferences :many
select e.id as entity_id, e.entity_name as entity_name, e.wbrn as resource_name, t.wbtn as type_name
from world.entities e inner join world.types t on e.type_id = t.id
where e.parent_id = $1;

-- name: GetEntityReference :one
select e.ID as entity_id, e.entity_name as entity_name, e.wbrn as resource_name, t.wbtn as type_name
from world.entities e inner join world.types t on e.type_id = t.id
where e.id = $1;

-- name: GetEntityReferenceByWBRN :one
select e.ID as entity_id, e.entity_name as entity_name, e.wbrn as resource_name, t.wbtn as type_name
from world.entities e inner join world.types t on e.type_id = t.id
where e.wbrn = $1;

-- name: CreateEntityAttribute :one
insert into world.entity_attributes
(entity_id, attribute_id, attribute_value)
values
($1, $2, $3)
returning *;
