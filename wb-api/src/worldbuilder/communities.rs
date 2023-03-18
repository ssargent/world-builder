use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Pool, Postgres};
use uuid::Uuid;

use crate::model::AppState;

use super::Manager;

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Community {
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

pub struct CommunityManager {
    db: Pool<Postgres>,
}

impl CommunityManager {
    pub fn new(db: Pool<Postgres>) -> CommunityManager {
        CommunityManager { db: db }
    }
}

#[async_trait]
impl super::Manager<Community> for CommunityManager {
    async fn create(&self, entity: Community) -> Result<Community, super::WBError> {
        todo!()
    }

    async fn get_by_id(&self, id: Uuid) -> Result<Community, super::WBError> {
        todo!()
    }

    async fn get_by_wbrn(&self, wbn: String) -> Result<Community, super::WBError> {
        todo!()
    }

    async fn get_all(
        &self,
        skip: i32,
        take: i32,
    ) -> Result<super::PagedSet<Community>, super::WBError> {
        todo!()
    }
}
