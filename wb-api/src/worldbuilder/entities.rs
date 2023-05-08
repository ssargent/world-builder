use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;

#[derive(Debug, FromRow, Deserialize, Serialize)]
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
    pub attributes: Vec<AttributeDefinition>,
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
pub struct AttributeDefinition {
    pub id: Uuid,
    pub wbatn: String,
    pub attribute_name: String,
    pub label: String,
    pub data_type: String,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
pub struct EntityAttribute {
    pub id: Uuid,
    pub entity_id: Uuid,
    pub attribute_id: Uuid,
    pub attribute_value: String,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

#[derive(Debug, FromRow, Deserialize, Serialize)]
pub struct TypeAttribute {
    pub type_id: Uuid,
    pub attribute_id: Uuid,
    pub ordinal: i32,
    pub is_required: bool,
}
