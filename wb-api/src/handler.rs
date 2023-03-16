use crate::{
    model::{AppState},
    response::{GenericResponse},
    schema::{FilterOptions},
    worldbuilder::{Country, Region, Community },
};
use actix_web::{get, web, HttpResponse, Responder}; 
use serde_json::json; 

#[get("/healthchecker")]
async fn health_checker_handler() -> impl Responder {
    const MESSAGE: &str = "wb-api (World Builder API)";

    let response_json = &GenericResponse{
        status: "success".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}
 
#[get("/countries")]
pub async fn country_list_handler( opts: web::Query<FilterOptions>, data: web::Data<AppState>) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let query_result = sqlx::query_as!(
        Country,
        "select * from world.countries order by id limit $1 offset $2",
        limit as i32,
        offset as i32
    )
    .fetch_all(&data.db)
    .await;

    if query_result.is_err() {
        let message = "Something bad happened while fetching all countries";
        return HttpResponse::InternalServerError().json(json!({"status": "error", "message": message}));
    }

    let countries = query_result.unwrap();

    let json_response = serde_json::json!({
        "status":"success",
        "results": countries.len(),
        "countries": countries
    });

    HttpResponse::Ok().json(json_response)
}

#[get("/regions")]
pub async fn region_list_handler( opts: web::Query<FilterOptions>, data: web::Data<AppState>) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let query_result = sqlx::query_as!(
        Region,
        r#"select id, wbrn, region_name, country_id, created_at, updated_at from world.regions order by id limit $1 offset $2"#, 
        limit as i32,
        offset as i32
    )
    .fetch_all(&data.db)
    .await;

    if query_result.is_err() {
        let message = "Something bad happened while fetching all regions";
        return HttpResponse::InternalServerError().json(json!({"status": "error", "message": message}));
    }

    let regions = query_result.unwrap();

    let json_response = serde_json::json!({
        "status":"success",
        "results": regions.len(),
        "regions": regions
    });

    HttpResponse::Ok().json(json_response)
}

#[get("/communities")]
pub async fn community_list_handler( opts: web::Query<FilterOptions>, data: web::Data<AppState>) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let query_result = sqlx::query_as!(
        Community,
        r#"select * from world.communities order by id limit $1 offset $2"#,
        limit as i32,
        offset as i32
    )
    .fetch_all(&data.db)
    .await;

    if query_result.is_err() {
        let message = "Something bad happened while fetching all communities";
        return HttpResponse::InternalServerError().json(json!({"status": "error", "message": message}));
    }

    let communities = query_result.unwrap();

    let json_response = serde_json::json!({
        "status":"success",
        "results": communities.len(),
        "communities": communities
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