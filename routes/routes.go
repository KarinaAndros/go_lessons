package routes

import (
	"backend/handlers"
	// "backend/middleware"
	"net/http"
)

func Routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/user", handlers.LoginHandler)

	// mux.HandleFunc("/api/profile", middleware.AuthMiddleware(handlers.ProfileHandler))
	return mux
}
