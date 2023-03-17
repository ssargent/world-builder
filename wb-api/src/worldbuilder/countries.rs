use actix_web::web;
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use sqlx::{FromRow, Pool, Postgres};
use uuid::Uuid;

use crate::model::AppState;

use super::PagedSet;

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Country {
    pub id: Uuid,
    pub wbrn: String,
    #[serde(rename = "countryName")]
    pub country_name: String,
    #[serde(rename = "createdAt")]
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    #[serde(rename = "updatedAt")]
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

pub struct CountryManager {}

impl CountryManager {
    pub fn new() -> CountryManager {
        CountryManager {}
    }

    pub fn hello(&self) -> String {
        let s = "hello world";
        s.to_string()
    }
}

#[async_trait]
impl super::Manager<Country> for CountryManager {
    async fn create(
        &self,
        data: web::Data<AppState>,
        entity: Country,
    ) -> Result<Country, super::WBError> {
        todo!()
    }

    async fn get_by_id(
        &self,
        data: web::Data<AppState>,
        id: Uuid,
    ) -> Result<Country, super::WBError> {
        todo!()
    }

    async fn get_all(
        &self,
        data: web::Data<AppState>,
        skip: i32,
        take: i32,
    ) -> Result<super::PagedSet<Country>, super::WBError> {
        let query_result = sqlx::query_as!(
            Country,
            "select * from world.countries order by id limit $1 offset $2",
            take as i32,
            skip as i32
        )
        .fetch_all(&data.db)
        .await;

        let countries = query_result.unwrap();

        Ok(PagedSet {
            results: countries.len(),
            items: countries,
            skip: skip,
            take: take,
        })
    }
}
