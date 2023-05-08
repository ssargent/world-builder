use crate::{db, errors::Error, worldbuilder::entities};

use super::Repository;
use log::error;
use sqlx;

impl Repository {
    pub async fn create_entity<'c, C: db::Queryer<'c>>(
        &self,
        db: C,
        entity: &entities::Entity,
    ) -> Result<(), Error> {
        const QUERY: &str = "insert into world.entities
        (type_id, parent_id, wbrn, entity_name, entity_description)
        values ( $1, $2, $3, $4, $5 )";

        match sqlx::query(QUERY)
            .bind(entity.type_id)
            .bind(entity.parent_id)
            .bind(&entity.wbrn)
            .bind(&entity.entity_name)
            .bind(&entity.entity_description)
            .execute(db)
            .await
        {
            Err(err) => {
                error!("worldbuilder.create_entity: Inserting entity: {}", &err);
                Err(err.into())
            }
            Ok(_) => Ok(()),
        }
    }
}
