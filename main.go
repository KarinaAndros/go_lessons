package main

import (
	"backend/database"
	"backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/gothic"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("Предупреждение: .env не найден")
    }
    goth.UseProviders(
        google.New(
            os.Getenv("GOOGLE_CLIENT_ID"),
            os.Getenv("GOOGLE_CLIENT_SECRET"),
            os.Getenv("CALLBACK_URL"),
            "email", "profile",
        ),
    )
}

func main(){
	//инициализация хранилища сессий
	key := "secret-session-key" 
  store := sessions.NewCookieStore([]byte(key))
  gothic.Store = store
	//инциализация бд
	database.InitDB()
	defer database.DB.Close()
	//получение маршрутов
	router := routes.Routes()
	http.ListenAndServe(":8080", router)
}
