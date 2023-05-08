use actix_web::{
    http::{header::ContentType, StatusCode},
    HttpResponse, ResponseError,
};
use serde::Serialize;

use crate::errors;

#[derive(Serialize)]
pub struct GenericResponse {
    pub status: String,
    pub message: String,
}

impl ResponseError for crate::errors::Error {
    fn status_code(&self) -> StatusCode {
        match *self {
            errors::Error::AlreadyExists(_) => StatusCode::CONFLICT,
            errors::Error::AuthenticationRequired => StatusCode::UNAUTHORIZED,
            errors::Error::Internal(_) => StatusCode::INTERNAL_SERVER_ERROR,
            errors::Error::NotFound(_) => StatusCode::NOT_FOUND,
            errors::Error::PermissionDenied(_) => StatusCode::FORBIDDEN,
            errors::Error::InvalidArgument(_) => StatusCode::BAD_REQUEST,
        }
    }

    fn error_response(&self) -> HttpResponse {
        HttpResponse::build(self.status_code())
            .insert_header(ContentType::json())
            .json(GenericResponse {
                status: String::from("error"),
                message: self.to_string(),
            })
    }
}
