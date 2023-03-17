use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;

use crate::model::AppState;

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Region {
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

pub struct RegionManager {}

#[async_trait]
impl super::Manager<Region> for RegionManager {
    async fn create(
        &self,
        data: web::Data<AppState>,
        entity: Region,
    ) -> Result<Region, super::WBError> {
        todo!()
    }

    async fn get_by_id(
        &self,
        data: web::Data<AppState>,
        id: Uuid,
    ) -> Result<Region, super::WBError> {
        todo!()
    }

    async fn get_all(
        &self,
        data: web::Data<AppState>,
        skip: i32,
        take: i32,
    ) -> Result<super::PagedSet<Region>, super::WBError> {
        todo!()
    }
}
