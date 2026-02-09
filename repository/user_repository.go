package repository

import (
	"backend/database"
	"backend/models"
)

func UpdateUser(email string, newData models.User) error {
	query := "UPDATE users SET name = COALESCE(?, name), surname = COALESCE(?, surname), avatar = COALESCE(?, avatar) WHERE email = ?"
	_, err := database.DB.Exec(query, newData.Name, newData.Surname, newData.Avatar, email)
	return err
}
