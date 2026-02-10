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
	mux.HandleFunc("/api/users", middleware.AuthMiddleware(middleware.AdminMiddleware(handlers.GetUsers)))
	mux.HandleFunc("/api/user", middleware.AuthMiddleware(handlers.GetUserData))
	mux.HandleFunc("/api/user/edit", middleware.AuthMiddleware(handlers.EditData))
	return mux
}
