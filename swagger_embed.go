package embed

import "embed"

// Embed Swagger JSON files
//
//go:embed docs/swagger/**
var SwaggerJSON embed.FS

// Embed Swagger UI assets
//
//go:embed swagger-ui/**
var SwaggerUI embed.FS
