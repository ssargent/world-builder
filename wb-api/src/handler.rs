use std::collections::HashMap;

use crate::{
    errors::{self, Error},
    model::{AppState, CreateEntity, Entity},
    response::GenericResponse,
    worldbuilder::service,
};
use actix_web::{
    get, post, web, web::Json, CustomizeResponder, HttpRequest, HttpResponse, Responder,
};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[get("/healthchecker")]
async fn health_checker_handler() -> impl Responder {
    const MESSAGE: &str = "wb-api (World Builder API)";

    let response_json = &GenericResponse {
        status: "success".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}

#[get("entities/{id}")]
async fn get_entity_handler(
    state: web::Data<AppState>,
    id: web::Path<Uuid>,
) -> Result<HttpResponse, errors::Error> {
    let res = state.service.get_entity(id.into_inner()).await?;
    Ok(HttpResponse::Ok().json(res))
}

#[get("entities/rn/{name}")]
async fn get_entity_by_wbrn_handler(
    state: web::Data<AppState>,
    wbrn: web::Path<String>,
) -> Result<HttpResponse, errors::Error> {
    print!("wbrn := {}\n", wbrn);
    let res = state.service.get_entity_by_wbrn(wbrn.into_inner()).await?;
    Ok(HttpResponse::Ok().json(res))
}

#[post("entities")]
async fn create_entity_handler(
    state: web::Data<AppState>,
    input: Json<CreateEntity>,
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
        Err(err) => match err {
            Error::NotFound(msg) => HttpResponse::NotFound().json(GenericResponse {
                status: "not found".to_string(),
                message: msg,
            }),
            Error::Internal(msg) => HttpResponse::InternalServerError().json(GenericResponse {
                status: "internal error".to_string(),
                message: msg,
            }),
            _ => HttpResponse::InternalServerError().json(GenericResponse {
                status: "internal error".to_string(),
                message: "unknown error".to_string(),
            }),
        },
    }
}

pub fn config(conf: &mut web::ServiceConfig) {
    let scope = web::scope("/api")
        .service(health_checker_handler)
        .service(create_entity_handler)
        .service(get_entity_handler)
        .service(get_entity_by_wbrn_handler);

    conf.service(scope);
}
