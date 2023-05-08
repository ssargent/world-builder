use uuid::Uuid;

use crate::{errors, worldbuilder::entities::Entity};

use super::Service;

impl Service {
    /// .
    ///
    /// # Errors
    ///
    /// This function will return an error if .
    pub async fn get_entity(&self, id: Uuid) -> Result<Entity, errors::Error> {
        let entity = self.repo.get_entity(&self.db, &id).await?;
        Ok(entity)
    }
}
