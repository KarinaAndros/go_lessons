package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(){
	var err error
	dns:=os.Getenv("DB_USER_NAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DB_NAME")
	DB, err = sql.Open("mysql", dns)
	if err != nil {log.Fatal("Ошибка в подключении", err)}
	err = DB.Ping()
	if err != nil {log.Fatal("Ошибка в подключении", err)}
	fmt.Println("Успешное подключение")
}
