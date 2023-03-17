use std::sync::Arc;

use serde::{Deserialize};
use sqlx::{Postgres, Pool};

use crate::worldbuilder::CountryManager;

// todo: Add RegionManager, CommunityManager and eventually EntityManager
pub struct AppState {
    pub db: Pool<Postgres>,
    pub country_manager: CountryManager,
}

impl AppState {
    pub fn init(pool: Pool<Postgres>) -> AppState {

        let cm = CountryManager::new(pool.clone());
        AppState {
            db:  pool,
            country_manager: cm,
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