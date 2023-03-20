use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Pool, Postgres};
use uuid::Uuid;

use crate::model::AppState;

use super::PagedSet;

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Entity {
    pub id: Uuid,
    pub type_id: Uuid,
    pub parent_id: Uuid,
    pub wbrn: String,
    #[serde(rename = "entityName")]
    pub entity_name: String,
    #[serde(rename = "entityDescription")]
    pub entity_description: String,
    pub notes: String,
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

impl EntityManager for EntityManagerImpl {
    fn get_by_id(&self, id: Uuid) -> Result<Entity, WBError> {
        todo!()
    }

    fn get_by_parent(&self, parent: Uuid) -> Result<Vec<Entity>, WBError> {
        todo!()
    }

    fn get_by_wbrn(&self, wbrn: String) -> Result<Vec<Entity>, WBError> {
        todo!()
    }

    fn get_by_type(&self, type_id: Uuid) -> Result<Vec<Entity>, WBError> {
        todo!()
    }

    fn create(&self, ent: Entity) -> Result<Entity, WBError> {
        todo!()
    }

    fn update(&self, ent: Entity) -> Result<Entity, WBError> {
        todo!()
    }

    fn delete(&self, ent: Entity) -> Result<(), WBError> {
        todo!()
    }
}
