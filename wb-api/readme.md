# wb-api 

World Builder API 

## Prerequisites

1. Postgres installed and running (can be via docker)
2. .env file created (copy .env.sample, change required info)

## Quick Start

```
you@computer: ~/source/github/ssargent/world-builder/wb-api$ make watch
```

## Types 

The world is made up of entities, and entities have types (Person, Place, Boat). 

There are primary entities that exist and are locateable on their own: `The Capital City` and there are secondary entities that are only locatable in the context of a primary entity: `Mike's Tavern in The Capital City`

`wb:[country]:[region]:[city]:[typename]:[entity-id]`

### Type Attributes

Types have attributes, a Tavern is a place to get food & drink, spend the night and a place to meet others.

### Countries 

Everything exists in a country, so in the context of the country you can identify the capital city

`wb:tomiland:capital:capital-city` would refer to the Capital City, in the Capital District in the country Tomiland.

### People

People belong to a country and live in a city or region, for example Frank Smith could live in the Capital City of Tomiland.

`wb:tomiland:capital:capital-city:people:frank-smith`

### Buildings

Buildings are also in a city, which exists in a district in a country.  

`wb:tomiland:capital:capital-city:building:mikes-tavern`

### Ships

Ships can be docked, in a city, in a district in a country

`wb:tomiland:capital:capital-city:ship:the-clement`

but ships also can be located outside of the geographic region

`wb:*:*:*:ship:the-clement`
