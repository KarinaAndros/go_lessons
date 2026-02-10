package models

import "errors"

type User struct{
	ID int `json:"id"`
	Name *string `json:"name"`
	Surname *string `json:"surname"`
	Password string `json:"password"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Role_id int `json:"role_id"`
}

func (u *User) Validate() error{
	if u.Name != nil && *u.Name == ""{
		return errors.New("имя не может быть пустым")
	}
	return nil
}
