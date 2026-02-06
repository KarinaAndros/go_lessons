package routes

import (
	"backend/handlers"
	"backend/middleware"
	"net/http"
)

func Routes() *http.ServeMux{
	mux := http.NewServeMux()
	//login routes
	mux.HandleFunc("/api/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/login", handlers.LoginHandler)

	//auth user routes
	mux.HandleFunc("/api/users", middleware.AuthMiddleware(handlers.GetUsers))
	mux.HandleFunc("/api/user", middleware.AuthMiddleware(handlers.GetUserData))
	return mux
}
