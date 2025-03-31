package api

import (
	"log"
	"net/http"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/handlers"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/services"
	"github.com/gorilla/mux"
)

// APIServer represents the API server
type APIServer struct {
	addr  string
	route handlers.User
}

// NewAPIServer creates a new APIServer instance
func NewAPIServer(addr string, route handlers.User) *APIServer {
	return &APIServer{
		addr:  addr,
		route: route,
	}
}

// Serve starts the API server and listens for incoming requests
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Register routes for different services
	userService := services.NewUserService(s.route)
	userService.RegisterRoutes(subrouter)

	adminService := services.NewAdminService(s.route)
	adminService.RegisterRoutes(subrouter)

	log.Println("Starting the API server at", s.addr)

	// Start the server and listen for incoming requests
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
