
use std::{collections::HashSet, fmt::Debug, sync::Arc};
use  uuid::Uuid;

use crate::db;

use super::repository::Repository;

mod create_entity;

#[derive(Debug, Clone)]
pub struct CreateEntityInput {
    pub wbtn: String,
    pub parent: String,
    pub entity_name: String,
    pub entity_description: String,
}

#[derive(Debug)]
pub struct Service {
    repo: Repository,
    db: db::DB,
}

impl Service {
    pub fn new(db: db::DB) -> Service {
        let repo = Repository::new();
        Service {
            db,
            repo,
        }
    }
}