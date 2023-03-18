-- Add up migration script here
insert into world.entities (
        id,
        type_id,
        parent_id,
        wbrn,
        entity_name,
        entity_description
    )
select uuid_generate_v5(
        uuid_ns_url(),
        '/v1/countries/' || lower(country_name)
    ),
    uuid_generate_v5(uuid_ns_url(), '/v1/types/country'),
    uuid_generate_v5(
        uuid_ns_url(),
        '/v1/worlds/ktos'
    ),
    'wb:ktos:' || lower(country_name),
    country_name,
    'Country of ' || country_name
    from world.countries;
-- add regions to entities by joining back to country
insert into world.entities (
        id,
        type_id,
        parent_id,
        wbrn,
        entity_name,
        entity_description
    )
select uuid_generate_v5(
        uuid_ns_url(),
        '/v1/regions/' || lower(region_name)
    ),
    uuid_generate_v5(uuid_ns_url(), '/v1/types/region'),
    uuid_generate_v5(
        uuid_ns_url(),
        '/v1/countries/' || lower(country_name)
    ),
    'wb:ktos:' || lower(c.country_name) || ':' || replace(lower(r.region_name), ' ', '-'),
    region_name,
    region_name || ' Region'
from world.regions r
    inner join world.countries c on c.id = r.country_id;


-- add communities to entities by joining back to region and country
insert into world.entities (
        id,
        type_id,
        parent_id,
        wbrn,
        entity_name,
        entity_description
    )
select uuid_generate_v5(
        uuid_ns_url(),
        '/v1/communities/' || replace(lower(m.community_name), ' ', '-')
    ),
    uuid_generate_v5(uuid_ns_url(), '/v1/types/community'),
    uuid_generate_v5(
        uuid_ns_url(),
        '/v1/regions/' || replace(lower(r.region_name), ' ', '-')
    ),
    'wb:ktos:' || lower(c.country_name) || ':' || replace(lower(r.region_name), ' ', '-') || ':' || replace(lower(m.community_name), ' ', '-'),
    community_name,
    community_name || ' Community'
from world.regions r
    inner join world.countries c on c.id = r.country_id
    inner join world.communities m on r.id = m.region_id;