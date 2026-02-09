package handlers

import (
	"backend/database"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"database/sql"
	"fmt"
	"net/http"
)

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
	query := "SELECT id, name, surname, email, avatar FROM users WHERE email = ?"
	var user models.User
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Avatar)
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

//edit auth user data
func EditData(w http.ResponseWriter, r *http.Request){
	//проверяем метод
	if !utils.CheckMethod(r, w, http.MethodPost){return}
	//получаем email из context()
	email := utils.GetEmail(w, r)
	if email == ""{return}
	//декодируем данные из тела запроса
	var newData models.User
	if !utils.DecodeData(r, w, &newData){ return }
	//выполняем изменение данных
	err := repository.UpdateUser(email, newData)
	if !utils.CheckError(w, err, "Ошибка обновления данных", http.StatusInternalServerError){return}
	//возвращаем ответ
	utils.ReturnResponse(w, map[string]string{"message" : "данные успешно обновлены!"}, http.StatusOK)
}
