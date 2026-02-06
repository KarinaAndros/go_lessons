package models

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Password string `json:"password"`
	Email string `json:"email"`
	Avatar string `json:"avtar"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
