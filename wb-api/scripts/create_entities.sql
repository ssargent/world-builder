-- create country
insert into world.entities
(type_id, parent_id, wbrn, entity_name, entity_description)
values
(
  uuid_generate_v5(uuid_ns_url(), '/v1/types/country'), 
 '4d4a7f6f-b407-52bf-bf04-09157aa494fb',
 'wb:ktos:tropica', 
 'Tropica', 
 'The Oceanic Country of Tropica'
)

-- create region
insert into world.entities
(type_id, parent_id, wbrn, entity_name, entity_description)
values
(
  uuid_generate_v5(uuid_ns_url(), '/v1/types/region'), 
 'fee430b7-b73e-40d9-b3eb-7c484c020ef8',
 'wb:ktos:skarlaand:capital', 
 'Capital', 
 'Capital Region'
)

-- create city
insert into world.entities
(type_id, parent_id, wbrn, entity_name, entity_description)
values
(
  uuid_generate_v5(uuid_ns_url(), '/v1/types/community'), 
 '69e8c20c-c18b-477f-bf0f-7a76a2d24bde',
 'wb:ktos:skarlaand:northern-coastal:eschton', 
 'Eschton', 
 'Eschton City'
)

-- create district
insert into world.entities
(type_id, parent_id, wbrn, entity_name, entity_description)
values
(
  uuid_generate_v5(uuid_ns_url(), '/v1/types/district'), 
 'd52e77db-209e-4127-95e8-47f8dcdb85bb',
 'wb:ktos:skarlaand:northern-coastal:eschton:emerald-park', 
 'Emerald Park', 
 'Neighborhood of Emerald Park'
)