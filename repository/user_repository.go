package repository

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"log"
)

//update user data
func UpdateUser(email string, newData models.User) error {
	query := "UPDATE users SET name = COALESCE(?, name), surname = COALESCE(?, surname), avatar = COALESCE(?, avatar) WHERE email = ?"
	_, err := database.DB.Exec(query, newData.Name, newData.Surname, newData.Avatar, email)
	return err
}

//get users for surname
func GetUserBySurname(surname string) ([]models.User, error){
	pattern := "%" + surname + "%"
	query := "SELECT id, name, surname, avatar FROM users WHERE surname LIKE ?"
	rows, err := database.DB.Query(query, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Avatar)
		if err != nil{
			log.Println("Ошибка сканирования данных")
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
