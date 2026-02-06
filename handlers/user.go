package handlers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"database/sql"
	"fmt"
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

//get auth user data
func GetUserData(w http.ResponseWriter, r *http.Request){
	//проверяем метод
	if !utils.CheckMethod(r, w, http.MethodGet){return}
	//достаём email из context()
	email, ok := r.Context().Value("email").(string)
	if !ok {
		utils.CheckError(w, fmt.Errorf("context error"), "Ошибка авторизации", http.StatusUnauthorized)
    return
	}
	//используем email для доступа к данным авторизованного пользователя
	query := "SELECT id, name, email FROM users WHERE email = ?"
	var user models.User
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
				utils.CheckError(w, err, "Пользователь не найден", http.StatusNotFound)
				return
		}
		utils.CheckError(w, err, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}
	//возвращаем данные о пользователе
	utils.ReturnResponse(w, user, http.StatusOK)
}
