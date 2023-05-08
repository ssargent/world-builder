use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;

#[derive(Debug, FromRow, Deserialize, Serialize)]
#[allow(non_snake_case)]
pub struct Entity {
    pub id: Uuid,
    pub type_id: Uuid,

    pub parent_id: Uuid,
    pub wbrn: String,
    pub entity_name: String,
    pub entity_description: String,
    pub notes: Option<String>,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
pub struct Type {
    pub id: Uuid,
    pub parent_id: Uuid,
    pub wbtn: String,
    pub type_name: String,
    pub type_description: String,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}
