use std::{sync::Arc, collections::HashMap, future::Ready};

use crate::{
    response::GenericResponse, model::{AppState, CreateEntity}, worldbuilder::{service, self}, errors::{self, Error},
};
use actix_web::{get, web::{self, Json}, HttpResponse, Responder, post, ResponseError, http::StatusCode, HttpRequest};
use serde::{Serialize, Deserialize};
 
 
#[get("/healthchecker")]
async fn health_checker_handler() -> impl Responder {
    const MESSAGE: &str = "wb-api (World Builder API)";

    let response_json = &GenericResponse {
        status: "success".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}

#[post("entities")]
async fn create_entity_handler(
    state: web::Data<AppState>,
    input: Json<CreateEntity>
) -> impl Responder {
    let input = input.into_inner();
    let service_input = service::CreateEntityInput {
        wbtn: input.wbtn,
        parent: input.parent,
        entity_name: input.entity_name,
        entity_description: input.entity_description,
    };
     
    match state.service.create_entity(service_input).await {
        Ok(res) => HttpResponse::Ok().json(res),
        Err(err) => {
            match err {
                Error::NotFound(msg) => HttpResponse::NotFound().json(GenericResponse{ status: "not found".to_string(), message: msg}),
                Error::Internal(msg) => HttpResponse::InternalServerError().json(GenericResponse{ status: "internal error".to_string(), message: msg}),
                _ => HttpResponse::InternalServerError().json(GenericResponse{ status: "internal error".to_string(), message: "unknown error".to_string()}),
            }
        }
    }
}


pub fn config(conf: &mut web::ServiceConfig) {
    let scope = web::scope("/api")
        .service(health_checker_handler)
        .service(create_entity_handler);

    conf.service(scope);
}
