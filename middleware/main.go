package middleware

import (
	"backend/utils"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//проверка на наличие заголовка
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			if !utils.CheckError(w, nil ,"Отсутствует заголовок авторизации", http.StatusUnauthorized){ return }
		}
		//получаем токен из заголовка
		parts := strings.Split(authHeader, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			if !utils.CheckError(w, nil ,"Неверный формат заголовка", http.StatusUnauthorized){ return }
		}
		//провека токена
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !utils.CheckError(w, err ,"Невалидный токен", http.StatusUnauthorized){ return }
		//получение email из токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := context.WithValue(r.Context(), "email", claims["email"])
			next(w, r.WithContext(ctx))
		} else {
			next(w, r)
		}
	}
}
