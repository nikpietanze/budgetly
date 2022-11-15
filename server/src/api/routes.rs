use super::messages;
use super::accounts;
use actix_web::{web, Scope};

pub fn routes() -> Scope {
    web::scope("/api")
        .service(messages::routes())
        .service(accounts::routes())
}
