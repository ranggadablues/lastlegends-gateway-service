package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPHandler() http.Handler {
	mux := mux.NewRouter()

	// Mount service-specific routes
	gateway := NewGateway()
	RegisterUserRoutes(mux, gateway.User)
	RegisterProductRoutes(mux, gateway.Product)

	registerSwaggerRoutes(mux)
	return mux
}

func registerSwaggerRoutes(mux *mux.Router) {
	// Serve Swagger JSONs (generated from protos)
	docsfile := http.Dir("../docs/swagger")
	mux.PathPrefix("/swagger/").Handler(
		http.StripPrefix("/swagger/",
			http.FileServer(docsfile),
		),
	)

	// Serve Swagger UI
	mux.PathPrefix("/docs/").Handler(
		http.StripPrefix("/docs/",
			http.FileServer(http.Dir("../swagger-ui")),
		),
	)
}
