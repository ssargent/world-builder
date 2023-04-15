mod communities;
mod countries;
mod entities;
mod regions;

use actix_web::web;
use async_trait::async_trait;
pub use communities::{Community, CommunityManager};
pub use countries::{Country, CountryManager};
pub use entities::{Entity, EntityManager, EntityManagerImpl};
pub use regions::{Region, RegionManager};

use serde::{Deserialize, Serialize};
use thiserror::Error;
use uuid::Uuid;

use crate::model::AppState;

#[async_trait]
pub trait Manager<T> {
    async fn create(&self, entity: T) -> Result<T, WBError>;
    async fn get_by_id(&self, id: Uuid) -> Result<T, WBError>;
    async fn get_all(&self, skip: i32, take: i32) -> Result<PagedSet<T>, WBError>;
    async fn get_by_wbrn(&self, wbn: String) -> Result<T, WBError>;
}

#[derive(Serialize, Debug, Deserialize)]
pub struct PagedSet<T> {
    pub items: Vec<T>,
    //  total_count: usize,
    pub results: usize,
    pub skip: i32,
    pub take: i32,
}

// todo: flesh this out a bit more.
#[derive(Error, Debug)]
pub enum WBError {
    #[error(transparent)]
    DatabaseError(#[from] sqlx::Error),
    #[error("a server error occurred: {0}.")]
    ServerError(String),
    #[error("an unknown error occurred.")]
    Unknown,
}
