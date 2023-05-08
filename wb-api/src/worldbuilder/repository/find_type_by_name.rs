use crate::{db, errors::worldbuilder::Error, worldbuilder::entities};

use super::Repository;
use log::error;
use sqlx;

impl Repository {
    pub async fn find_type_by_name<'c, C: db::Queryer<'c>>(
        &self,
        db: C,
        name: String,
    ) -> Result<entities::Type, Error> {
        const QUERY: &str = "select * from world.types where type_name = $1";

        match sqlx::query_as::<_, entities::Type>(QUERY)
            .bind(name)
            .fetch_optional(db)
            .await
        {
            Err(err) => {
                error!("worldbuilder.find_type_by_name: retrieving type: {}", &err);
                Err(err.into())
            }
            Ok(None) => Err(Error::TypeNotFound),
            Ok(Some(res)) => Ok(res),
        }
    }
}
