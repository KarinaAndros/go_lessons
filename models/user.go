package models

import (
	"time"
)

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Surname *string `json:"surname"`
	Password string `json:"password"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Lat *float64 `json:"lat"`
	Lng *float64 `json:"lng"`
	Role_id int `json:"role_id"`
	Created_at time.Time `json:"created_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}
