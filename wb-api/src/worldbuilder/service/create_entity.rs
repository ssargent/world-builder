use chrono::Utc;
use uuid::Uuid;

use crate::{errors, worldbuilder::entities::Entity};

use super::{CreateEntityInput, Service};

impl Service {
    pub async fn create_entity(&self, input: CreateEntityInput) -> Result<Entity, errors::Error> {
        let wb_type = self.repo.find_type_by_wbtn(&self.db, input.wbtn).await?;
        let parent = self
            .repo
            .find_entity_by_wbrn(&self.db, input.parent)
            .await?;
        let now = Utc::now();
        let slug = input.entity_name.to_lowercase().replace(" ", "-");

        let entity = Entity {
            id: Uuid::new_v4(),
            type_id: wb_type.id,
            parent_id: parent.id,
            wbrn: format!("{}:{}", &parent.wbrn, slug),
            entity_name: input.entity_name,
            entity_description: input.entity_description,
            notes: Some("".to_string()),
            created_at: Some(now),
            updated_at: Some(now),
        };

        self.repo.create_entity(&self.db, &entity).await?;
        Ok(entity)
    }
}
