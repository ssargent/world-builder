use crate::{db, errors::worldbuilder::Error, worldbuilder::entities};

use super::Repository;
use log::error;
use sqlx;

impl Repository {
    /// .
    ///
    /// # Errors
    ///
    /// This function will return an error if .
    pub async fn get_entity<'c, C: db::Queryer<'c>>(
        &self,
        db: C,
        id: &uuid::Uuid,
    ) -> Result<entities::Entity, Error> {
        const QUERY: &str = "select * from world.entities where id = $1";

        match sqlx::query_as::<_, entities::Entity>(QUERY)
            .bind(id)
            .fetch_optional(db)
            .await
        {
            Err(err) => {
                error!("worldbuilder.get_entity: retrieving entity: {}", &err);
                Err(err.into())
            }
            Ok(None) => Err(Error::EntityNotFound),
            Ok(Some(res)) => Ok(res),
        }
    }
}
