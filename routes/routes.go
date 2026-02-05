package routes

import (
	"backend/handlers"
	"net/http"
)

func Routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/user", handlers.LoginHandler)
	return mux
}
