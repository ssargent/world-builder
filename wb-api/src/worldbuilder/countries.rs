use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;

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
        CountryManager {  }
    }

    pub fn hello(&self) -> String {
        let s = "hello world";
        s.to_string()
    }
}
