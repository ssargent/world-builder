use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;

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
 