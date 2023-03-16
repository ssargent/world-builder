use crate::{
    model::{AppState, QueryOptions, Todo, UpdateTodoSchema, CountryModel, RegionModel, CommunityModel},
    response::{GenericResponse, SingleTodoResponse, TodoData, TodoListResponse},
    schema::{FilterOptions, CreateCountrySchema, UpdateCountrySchema},
};
use actix_web::{delete, get, patch, post, web, HttpResponse, Responder};
use chrono::prelude::*;
use serde_json::json;
use uuid::Uuid;

#[get("/healthchecker")]
async fn health_checker_handler() -> impl Responder {
    const MESSAGE: &str = "wb-api (World Builder API)";

    let response_json = &GenericResponse{
        status: "success".to_string(),
        message: MESSAGE.to_string(),
    };

    HttpResponse::Ok().json(response_json)
}

#[get("/todos")]
async fn todos_list_handler(opts: web::Query::<QueryOptions>, data: web::Data<AppState>) -> impl Responder {
    let todos = data.todo_db.lock().unwrap();

    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let totalTodos = todos.len();
    let todos: Vec<Todo> = todos.clone().into_iter().skip(offset).take(limit).collect();

    let json_response = TodoListResponse{
        status: "success".to_string(),
        total_results: totalTodos,
        results: todos.len(),
        todos,
    };

    HttpResponse::Ok().json(json_response)
}

#[post("/todos")]
async fn create_todo_handler(mut body: web::Json<Todo>, data: web::Data<AppState>) -> impl Responder {
    let mut vec = data.todo_db.lock().unwrap();

    let todo = vec.iter().find(|todo| todo.title == body.title);

    if todo.is_some() {
        let error_response = GenericResponse{
            status: "fail".to_string(),
            message: format!("Todo with title '{}' already exists", body.title),
        };

        return HttpResponse::Conflict().json(error_response);
    }

    let uuid_id = Uuid::new_v4();
    let datetime = Utc::now();

    body.id = Some(uuid_id);
    body.completed = Some(false);
    body.createdAt = Some(datetime);
    body.updatedAt = Some(datetime);

    let todo = body.to_owned();

    vec.push(body.into_inner());

    let json_response = SingleTodoResponse{
        status: "success".to_string(),
        data: TodoData{ todo },
    };

    HttpResponse::Ok().json(json_response)
}

#[get("/countries")]
pub async fn country_list_handler( opts: web::Query<FilterOptions>, data: web::Data<AppState>) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let query_result = sqlx::query_as!(
        CountryModel,
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
        RegionModel,
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
        CommunityModel,
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
        .service(todos_list_handler)
        .service(create_todo_handler)
        .service(country_list_handler)
        .service(region_list_handler)
        .service(community_list_handler);

    conf.service(scope);
}