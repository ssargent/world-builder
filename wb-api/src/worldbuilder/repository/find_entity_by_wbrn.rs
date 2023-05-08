use crate::{db, errors::worldbuilder::Error, worldbuilder::entities};

use super::Repository;
use log::error;
use sqlx;

impl Repository {
    pub async fn find_entity_by_wbrn<'c, C: db::Queryer<'c>>(
        &self,
        db: C,
        wbrn: String,
    ) -> Result<entities::Entity, Error> {
        const QUERY: &str = "select * from world.entities where wbrn = $1";

        match sqlx::query_as::<_, entities::Entity>(QUERY)
            .bind(wbrn)
            .fetch_optional(db)
            .await
        {
            Err(err) => {
                error!(
                    "worldbuilder.find_entity_by_wbrn: retrieving entity: {}",
                    &err
                );
                Err(err.into())
            }
            Ok(None) => Err(Error::EntityNotFound),
            Ok(Some(res)) => Ok(res),
        }
    }
}
