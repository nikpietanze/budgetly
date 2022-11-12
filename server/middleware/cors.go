package middleware

import (
    "github.com/rs/cors"
)

func Cors() *cors.Cors {
    return cors.New(cors.Options{
        AllowedOrigins: []string{"*", "http://localhost:3000"},
        AllowCredentials: true,
    })
}

