
#[derive(Debug)]
pub enum Error {
    Internal,
    NotFound,
    EntityNotFound,
    TypeNotFound
}

impl std::convert::From<Error> for crate::errors::Error {
    fn from(err: Error) -> Self {
        match err {
            // Types
            Error::TypeNotFound => crate::errors::Error::NotFound(String::from("type not found")),
            // Entities
            Error::EntityNotFound => crate::errors::Error::NotFound(String::from("entity not found.")),
            Error::NotFound => crate::errors::Error::NotFound(String::new()),
            Error::Internal =>  crate::errors::Error::Internal(String::new()),
        }
    }
}

impl std::convert::From<sqlx::Error> for Error {
    fn from(err: sqlx::Error) -> Self {
        match err {
            // Not found error should be catched manually
            _ => Error::Internal,
        }
    }
}