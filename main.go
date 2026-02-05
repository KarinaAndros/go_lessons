package main

import (
	"backend/database"
	"backend/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main(){
	//загрузка .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	//инциализация бд
	database.InitDB()
	defer database.DB.Close()
	//получение маршрутов
	router := routes.Routes()
	http.ListenAndServe(":8080", router)
}
