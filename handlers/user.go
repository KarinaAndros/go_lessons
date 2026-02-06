package handlers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"log"
	"net/http"
)

//get all users
func GetUsers( w http.ResponseWriter, r *http.Request){
	//проверка метода
	if !utils.CheckMethod(r, w, http.MethodGet) {return}
	//отправка запроса
	query := "SELECT id, name, email FROM users"
	rows, err := database.DB.Query(query)
	if !utils.CheckError(w, err, "Ошибка получения пользователей", http.StatusInternalServerError){return}
	defer rows.Close()
	//создание слайса для хранения списка с пользователями
	var users []models.User
	//перебор строк из базы
	for rows.Next(){
		var user models.User
		err:= rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil{
			log.Println("Ошибка сканирования данных пользователя: ", err)
			continue
		}
		users = append(users, user)
	}
	//возврат ответа
	utils.ReturnResponse(w, users, http.StatusOK)
}

