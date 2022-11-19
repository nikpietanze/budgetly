use super::types::{Message, Metadata};
use crate::{extractors::Claims, types::ErrorMessage};
use actix_web::{get, web, Either, HttpResponse, Responder};
use std::collections::HashSet;

#[get("/admin")]
pub async fn admin(claims: Claims) -> Either<impl Responder, HttpResponse> {
    if claims.validate_permissions(&HashSet::from(["read:admin-messages".to_string()])) {
        Either::Left(web::Json(Message {
            metadata: Metadata {
                api: "api_actix-web_rust_hello-world".to_string(),
                branch: "basic-role-based-access-control".to_string(),
            },
            text: "This is only viewable to admins - Messages".to_string(),
        }))
    } else {
        Either::Right(HttpResponse::Forbidden().json(ErrorMessage {
            error: Some("insufficient_permissions".to_string()),
            error_description: Some("Requires read:admin-messages".to_string()),
            message: "Permission denied".to_string(),
        }))
    }
}

#[get("/protected")]
pub async fn protected(_claims: Claims) -> impl Responder {
    web::Json(Message {
        metadata: Metadata {
            api: "api_actix-web_rust_hello-world".to_string(),
            branch: "basic-role-based-access-control".to_string(),
        },
        text: "This should be protected and only viewable to the owner - Messages".to_string(),
    })
}

#[get("/public")]
pub async fn public() -> impl Responder {
    web::Json(Message {
        metadata: Metadata {
            api: "api_actix-web_rust_hello-world".to_string(),
            branch: "basic-role-based-access-control".to_string(),
        },
        text: "This is public and is viewable by anyone - Messages".to_string(),
    })
}
