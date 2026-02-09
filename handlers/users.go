package handlers

import (
	"backend/database"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"log"
	"net/http"
)

//get all users
func GetUsers( w http.ResponseWriter, r *http.Request){
	//проверка метода
	if !utils.CheckMethod(r, w, http.MethodGet) {return}
	//отправка запроса
	query := "SELECT id, name, email, avatar FROM users"
	rows, err := database.DB.Query(query)
	if !utils.CheckError(w, err, "Ошибка получения пользователей", http.StatusInternalServerError){return}
	defer rows.Close()
	//создание слайса для хранения списка с пользователями
	var users []models.User
	//перебор строк из базы
	for rows.Next(){
		var user models.User
		err:= rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
		if err != nil{
			log.Println("Ошибка сканирования данных пользователя: ", err)
			continue
		}
		users = append(users, user)
	}
	//возврат ответа
	utils.ReturnResponse(w, users, http.StatusOK)
}

func SearchUsersHandler(w http.ResponseWriter, r *http.Request){
	//проверяем метод
	if !utils.CheckMethod(r, w, http.MethodGet){return}
	//достаём фамилию из URL
	surname := r.URL.Query().Get("surname")
	if surname == ""{
		utils.ReturnResponse(w, map[string]string{"message": "Параметр surname обязателен"}, http.StatusBadRequest)
		return
	}
	//получаем список с пользователями
	users, err := repository.GetUserBySurname(surname)
	if !utils.CheckError(w, err, "Ошибка при поиске пользователей", http.StatusInternalServerError) {
  	return
  }
	//возвращаем результат
	utils.ReturnResponse(w, users, http.StatusOK)
}
