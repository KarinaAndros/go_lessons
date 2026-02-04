package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(){
	var err error
	dns:="karina:@tcp(127.0.0.1:3306)/diplom"
	DB, err = sql.Open("mysql", dns)
	if err != nil {log.Fatal("Ошибка в подключении", err)}
	err = DB.Ping()
	if err != nil {log.Fatal("Ошибка в подключении", err)}
	fmt.Println("Успешное подключение")
}
