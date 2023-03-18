use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Pool, Postgres};
use uuid::Uuid;

use crate::model::AppState;

use super::PagedSet;

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

pub struct RegionManager {
    db: Pool<Postgres>,
}

impl RegionManager {
    pub fn new(db: Pool<Postgres>) -> RegionManager {
        RegionManager { db: db }
    }

    pub fn get_by_country(&self, id: Uuid) -> Result<PagedSet<Region>, super::WBError> {
        todo!()
    }
}

#[async_trait]
impl super::Manager<Region> for RegionManager {
    async fn create(&self, entity: Region) -> Result<Region, super::WBError> {
        todo!()
    }

    async fn get_by_id(&self, id: Uuid) -> Result<Region, super::WBError> {
        let region = match sqlx::query_as!(
            Region,
            r#"select id, wbrn, region_name, country_id, created_at, updated_at from world.regions where id = $1"#, 
            id
        )
        .fetch_one(&self.db)
        .await {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(region)
    }

    async fn get_by_wbrn(&self, wbn: String) -> Result<Region, super::WBError> {
        todo!()
    }

    async fn get_all(
        &self,
        skip: i32,
        take: i32,
    ) -> Result<super::PagedSet<Region>, super::WBError> {
        let query_result = sqlx::query_as!(
            Region,
            r#"select id, wbrn, region_name, country_id, created_at, updated_at from world.regions order by id limit $1 offset $2"#, 
            take as i32,
            skip as i32
        )
        .fetch_all(&self.db)
        .await;

        let regions = match query_result {
            Ok(r) => r,
            Err(err) => return Err(super::WBError::DatabaseError(err)),
        };

        Ok(PagedSet {
            results: regions.len(),
            items: regions,
            skip: skip,
            take: take,
        })
    }
}
