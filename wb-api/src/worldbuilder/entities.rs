use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Pool, Postgres};
use uuid::Uuid;

use super::{PagedSet, WBError};

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Entity {
    pub id: Uuid,
    #[serde(rename = "typeId")]
    pub type_id: Uuid,
    #[serde(rename = "parentId")]
    pub parent_id: Uuid,
    pub wbrn: String,
    #[serde(rename = "entityName")]
    pub entity_name: String,
    #[serde(rename = "entityDescription")]
    pub entity_description: String,
    pub notes: Option<String>,
    #[serde(rename = "createdAt")]
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    #[serde(rename = "updatedAt")]
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

#[async_trait]
pub trait EntityManager {
    async fn get_by_id(&self, id: Uuid) -> Result<Entity, WBError>;
    async fn get_by_parent(&self, parent: Uuid) -> Result<Vec<Entity>, WBError>;
    async fn get_by_wbrn(&self, wbrn: String) -> Result<Vec<Entity>, WBError>;
    async fn get_by_type(&self, type_id: Uuid) -> Result<Vec<Entity>, WBError>;
    async fn create(&self, ent: Entity) -> Result<Entity, WBError>;
    async fn update(&self, ent: Entity) -> Result<Entity, WBError>;
    async fn delete(&self, ent: Entity) -> Result<(), WBError>;
}

pub struct EntityManagerImpl {
    db: Pool<Postgres>,
}

impl EntityManagerImpl {
    pub fn new(db: Pool<Postgres>) -> EntityManagerImpl {
        EntityManagerImpl { db: db }
    }
}
/*

create table world.entities (
    id uuid not null default (uuid_generate_v4()),
    type_id uuid not null,
    parent_id uuid not null,
    wbrn text not null,
    entity_name text not null,
    entity_description text not null,
    notes text null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint pk_entities_id primary key (id),
    constraint fk_entities_types foreign key (type_id) references world.types (id),
    constraint fk_entities_entities foreign key (parent_id) references world.entities (id)
);

 */
#[async_trait]
impl EntityManager for EntityManagerImpl {
    async fn get_by_id(&self, id: Uuid) -> Result<Entity, WBError> {
        let entity = match sqlx::query_as!(
            Entity,
            r#"select id, type_id, parent_id, wbrn, entity_name, entity_description, notes, created_at, updated_at from world.entities where id = $1"#, 
            id
        )
        .fetch_one(&self.db)
        .await {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(entity)
    }

    async fn get_by_parent(&self, parent: Uuid) -> Result<Vec<Entity>, WBError> {
        let entities = match sqlx::query_as!(
            Entity,
            r#"select id, type_id, parent_id, wbrn, entity_name, entity_description, notes, created_at, updated_at from world.entities where parent_id = $1"#, 
            parent
        )
        .fetch_all(&self.db)
        .await {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(entities)
    }

    async fn get_by_wbrn(&self, wbrn: String) -> Result<Vec<Entity>, WBError> {
        let entities = match sqlx::query_as!(
            Entity,
            r#"select id, type_id, parent_id, wbrn, entity_name, entity_description, notes, created_at, updated_at from world.entities where wbrn similar to $1"#, 
            wbrn
        )
        .fetch_all(&self.db)
        .await {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(entities)
    }

    async fn get_by_type(&self, type_id: Uuid) -> Result<Vec<Entity>, WBError> {
        let entities = match sqlx::query_as!(
            Entity,
            r#"select id, type_id, parent_id, wbrn, entity_name, entity_description, notes, created_at, updated_at from world.entities where type_id = $1"#, 
            type_id
        )
        .fetch_all(&self.db)
        .await {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(entities)
    }

    async fn create(&self, ent: Entity) -> Result<Entity, WBError> {
        todo!()
    }

    async fn update(&self, ent: Entity) -> Result<Entity, WBError> {
        todo!()
    }

    async fn delete(&self, ent: Entity) -> Result<(), WBError> {
        todo!()
    }
}
