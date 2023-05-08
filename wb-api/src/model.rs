use std::sync::Arc;

use serde::{Deserialize, Serialize};

use crate::worldbuilder::service::Service;
use sqlx::{Pool, Postgres};
use uuid::Uuid;

// todo: Add RegionManager, CommunityManager and eventually EntityManager
pub struct AppState {
    pub db: Pool<Postgres>,
    pub service: Arc<Service>,
}

impl AppState {
    pub fn init(pool: Pool<Postgres>, service: Arc<Service>) -> AppState {
        AppState {
            db: pool,
            service: service,
        }
    }
}

#[derive(Debug, Deserialize)]
pub struct QueryOptions {
    pub page: Option<usize>,
    pub limit: Option<usize>,
}

#[allow(non_snake_case)]
#[derive(Debug, Deserialize)]
pub struct UpdateTodoSchema {
    pub title: Option<String>,
    pub content: Option<String>,
    pub completed: Option<bool>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateEntity {
    pub wbtn: String,
    pub parent: String,
    pub entity_name: String,
    pub entity_description: String,
}

#[derive(Debug, Deserialize, Serialize)]
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
