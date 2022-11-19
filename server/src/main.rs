mod api;
mod extractors;
mod middlewares;
mod types;

use actix_web::{App, HttpServer};
use dotenv::dotenv;
use sea_orm::{Database, DatabaseConnection};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    env_logger::init();
    let config = types::Config::default();

    let db_url = config.db_url;
    let host = config.host;
    let port = config.port;

    let conn = Database::connect(db_url).await.unwrap();

    let auth0_config = extractors::Auth0Config::default();
    HttpServer::new(move || {
        App::new()
            .app_data(auth0_config.clone())
            .wrap(middlewares::cors(&config.client_origin_url))
            .wrap(middlewares::err_handlers())
            .wrap(middlewares::security_headers())
            .wrap(middlewares::logger())
            .service(api::routes())
    })
    .bind((host, port))?
    .run()
    .await
}
