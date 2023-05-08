pub mod entities;
pub mod repository;
pub mod service;

use serde::{Deserialize, Serialize};
use thiserror::Error;

#[derive(Serialize, Debug, Deserialize)]
pub struct PagedSet<T> {
    pub items: Vec<T>,
    //  total_count: usize,
    pub results: usize,
    pub skip: i32,
    pub take: i32,
}

// todo: flesh this out a bit more.
#[derive(Error, Debug)]
pub enum WBError {
    #[error(transparent)]
    DatabaseError(#[from] sqlx::Error),
    // #[error("a server error occurred: {0}.")]
    // ServerError(String),
    // #[error("an unknown error occurred.")]
    // Unknown,
}
