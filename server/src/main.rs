mod api;
mod extractors;
mod middlewares;
mod types;

use actix_web::{App, HttpServer, web::Data};
use diesel::{
    prelude::*,
    r2d2::{self, ConnectionManager}
};
use dotenv::dotenv;

type DbPool = r2d2::Pool<ConnectionManager<PgConnection>>;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    env_logger::init();

    let config = types::Config::default();
    let auth0_config = extractors::Auth0Config::default();

    let db_url = config.database_url;
    let manager = ConnectionManager::<PgConnection>::new(db_url);
    let pool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create pool.");

    HttpServer::new(move || {
        App::new()
            .app_data(auth0_config.clone())
            .app_data(Data::new(pool.clone()))
            .wrap(middlewares::cors(&config.client_origin_url))
            .wrap(middlewares::err_handlers())
            .wrap(middlewares::security_headers())
            .wrap(middlewares::logger())
            .service(api::routes())
    })
    .bind((config.host, config.port))?
    .run()
    .await
}
