
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