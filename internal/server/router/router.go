package router

import (
	handlers "coin-keeper/internal/server/handlers/engine"
	"net/http"
)

type Router struct {
	handlers *handlers.HandlersKeeper
}

func NewRouter(handlers *handlers.HandlersKeeper) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (r *Router) SetupRoutes() {
	http.HandleFunc("POST /api/v1/users", r.handlers.HandleCreateUser)
	http.HandleFunc("GET /api/v1/users/{id}", r.handlers.HandleReadUser)
	http.HandleFunc("PUT /api/v1/users", r.handlers.HandleUpdateUser)
	http.HandleFunc("DELETE /api/v1/users/{id}", r.handlers.HandleDeleteUser)
}
