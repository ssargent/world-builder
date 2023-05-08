use crate::{db, errors::worldbuilder::Error, worldbuilder::entities};

use super::Repository;
use log::error;
use sqlx;

impl Repository {
    pub async fn find_type_by_wbtn<'c, C: db::Queryer<'c>>(
        &self,
        db: C,
        wbtn: String,
    ) -> Result<entities::Type, Error> {
        const QUERY: &str = "select * from world.types where wbtn = $1";

        match sqlx::query_as::<_, entities::Type>(QUERY)
            .bind(wbtn)
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
