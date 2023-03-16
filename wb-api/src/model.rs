use serde::{Deserialize};
use sqlx::{Postgres, Pool};


pub struct AppState {
    pub db: Pool<Postgres>,
}

impl AppState {
    pub fn init(pool: Pool<Postgres>) -> AppState {
        AppState {
            db:  pool,
        }
    }
}

#[derive(Debug, Deserialize)]
pub struct QueryOptions {
    pub page: Option<usize>,
    pub limit: Option<usize>,
}

#[allow(non_snake_case)]
#[derive(Debug, Deserialize)]
pub struct UpdateTodoSchema {
    pub title: Option<String>,
    pub content: Option<String>,
    pub completed: Option<bool>,
}