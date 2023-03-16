use chrono::prelude::*;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Postgres, Pool};
use uuid::Uuid;
use std::sync::{Arc, Mutex};

#[allow(non_snake_case)]
#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct Todo {
    pub id: Option<Uuid>,
    pub title: String,
    pub content: String,
    pub completed: Option<bool>,
    pub createdAt: Option<DateTime<Utc>>,
    pub updatedAt: Option<DateTime<Utc>>,
}

pub struct AppState {
    pub todo_db: Arc<Mutex<Vec<Todo>>>,
    pub db: Pool<Postgres>,
}

impl AppState {
    pub fn init(pool: Pool<Postgres>) -> AppState {
        AppState {
            todo_db: Arc::new(Mutex::new(Vec::new())),
            db:  pool,
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

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct CountryModel {
    pub id: Uuid,
    pub wbrn: String,
    #[serde(rename = "countryName")]
    pub country_name: String,
    #[serde(rename = "createdAt")]
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    #[serde(rename = "updatedAt")]
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,   
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct RegionModel {
    pub id: Uuid,
    pub wbrn: String,
    #[serde(rename = "regionName")]
    pub region_name: String,
    #[serde(rename = "countryId")]
    pub country_id: Option<Uuid>,
    #[serde(rename = "createdAt")]
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    #[serde(rename = "updatedAt")]
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,   
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct CommunityModel {
    pub id: Uuid,
    pub wbrn: String,
    #[serde(rename = "communityName")]
    pub community_name: String,
    #[serde(rename = "regionId")]
    pub region_id: Option<Uuid>,
    #[serde(rename = "createdAt")]
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    #[serde(rename = "updatedAt")]
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,   
}