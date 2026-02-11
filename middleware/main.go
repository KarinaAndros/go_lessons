package middleware

import (
	"backend/database"
	"backend/utils"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

//autirization middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//проверка на наличие заголовка
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			if !utils.CheckError(w, nil ,"Отсутствует заголовок авторизации", http.StatusUnauthorized){ return }
		}
		//получаем токен из заголовка
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			if !utils.CheckError(w, nil ,"Неверный формат заголовка", http.StatusUnauthorized){ return }
		}
		//провека токена
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !utils.CheckError(w, err ,"Невалидный токен", http.StatusUnauthorized){ return }
		//получение email из токена и передача в context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := context.WithValue(r.Context(), "email", claims["email"])
			next(w, r.WithContext(ctx))
		} else {
			next(w, r)
		}
	}
}

//middleware for admin
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//получаем email
		email := utils.GetEmail(w, r)
		if email == "" { return }
		//получаем роль пользователя
		var roleTitle string
		query := "SELECT r.title FROM roles r JOIN users u ON u.role_id = r.id WHERE u.email = ?"
		err := database.DB.QueryRow(query, email).Scan(&roleTitle)
		if !utils.CheckError(w, err, "Поступ запрещён", http.StatusForbidden){return}
		//проверяем строку
		if roleTitle != "admin" {
			utils.CheckError(w, nil, "У вас нет прав администратора", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
