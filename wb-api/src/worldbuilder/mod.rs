mod communities;
mod countries;
mod regions;

use actix_web::web;
use async_trait::async_trait;
pub use communities::Community;
pub use countries::{Country, CountryManager};
pub use regions::Region;

use serde::{Deserialize, Serialize};
use thiserror::Error;
use uuid::Uuid;

use crate::model::AppState;

#[async_trait]
pub trait Manager<T> {
    async fn create(&self, data: web::Data<AppState>, entity: T) -> Result<T, WBError>;
    async fn get_by_id(&self, data: web::Data<AppState>, id: Uuid) -> Result<T, WBError>;
    async fn get_all(
        &self,
        data: web::Data<AppState>,
        skip: i32,
        take: i32,
    ) -> Result<PagedSet<T>, WBError>;
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
    #[error("a server error occurred.")]
    ServerError,
    #[error("an unknown error occurred.")]
    Unknown,
}
