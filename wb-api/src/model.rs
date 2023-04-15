use serde::Deserialize;
use sqlx::{Pool, Postgres};

use crate::worldbuilder::{
    CommunityManager, CountryManager, EntityManager, EntityManagerImpl, RegionManager,
};

// todo: Add RegionManager, CommunityManager and eventually EntityManager
pub struct AppState {
    pub db: Pool<Postgres>,
    pub country_manager: CountryManager,
    pub region_manager: RegionManager,
    pub community_manager: CommunityManager,
    pub entity_manager: EntityManagerImpl,
}

impl AppState {
    pub fn init(pool: Pool<Postgres>) -> AppState {
        let cm = CountryManager::new(pool.clone());
        let comu = CommunityManager::new(pool.clone());
        let rm = RegionManager::new(pool.clone());
        let em: EntityManagerImpl = EntityManagerImpl::new(pool.clone());

        AppState {
            db: pool,
            country_manager: cm,
            community_manager: comu,
            region_manager: rm,
            entity_manager: em,
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
