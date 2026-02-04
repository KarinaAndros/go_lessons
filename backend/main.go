package main

import(
	"backend/database"
)

func main(){
	database.InitDB()
	defer database.DB.Close()
}
