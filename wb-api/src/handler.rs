use crate::{
    model::AppState,
    response::GenericResponse,
    schema::FilterOptions,
    worldbuilder::{Community, Country, CountryManager, Manager, Region},
};
use actix_web::{get, web, HttpResponse, Responder};
use serde_json::json;

#[get("/healthchecker")]
async fn health_checker_handler() -> impl Responder {
    const MESSAGE: &str = "wb-api (World Builder API)";

    let response_json = &GenericResponse {
        status: "success".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}

#[get("/countries")]
pub async fn country_list_handler(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let countries = match data
        .country_manager
        .get_all(offset as i32, limit as i32)
        .await
    {
        Ok(c) => c,
        Err(err) => {
            return HttpResponse::InternalServerError()
                .json(json!({"status": "error", "message": err.to_string()}))
        }
    };

    let json_response = serde_json::json!({
        "status":"success",
        "countries": countries
    });

    HttpResponse::Ok().json(json_response)
}

#[get("/regions")]
pub async fn region_list_handler(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let regions = match data
        .region_manager
        .get_all(offset as i32, limit as i32)
        .await
    {
        Ok(r) => r,
        Err(err) => {
            return HttpResponse::InternalServerError()
                .json(json!({"status": "error", "message": err.to_string()}))
        }
    };

    let json_response = serde_json::json!({
        "status":"success",
        "regions": regions
    });

    HttpResponse::Ok().json(json_response)
}

#[get("/communities")]
pub async fn community_list_handler(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let communities = match data
        .community_manager
        .get_all(offset as i32, limit as i32)
        .await
    {
        Ok(r) => r,
        Err(err) => {
            return HttpResponse::InternalServerError()
                .json(json!({"status": "error", "message": err.to_string()}))
        }
    };

    let json_response = serde_json::json!({
        "status":"success",
        "communities": communities,
    });

    HttpResponse::Ok().json(json_response)
}

pub fn config(conf: &mut web::ServiceConfig) {
    let scope = web::scope("/api")
        .service(health_checker_handler)
        .service(country_list_handler)
        .service(region_list_handler)
        .service(community_list_handler);

    conf.service(scope);
}
