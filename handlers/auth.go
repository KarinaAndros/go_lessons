package handlers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//generate token for user
func GenerateToken(email string)(string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour*168).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Ошибка генерации токена", err)
		return "", err
	}
	return tokenString, nil
}

//registration
func RegisterHandler(w http.ResponseWriter, r *http.Request){
	//разрешение на использование только POST запроса
	if !utils.CheckMethod(r, w, http.MethodPost){ return }
	//декодирование JSON из тела запроса
	var user models.User
	if !utils.DecodeData(r, w, &user){ return }
	//хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if !utils.CheckError(w, err, "Ошибка хэширования пароля", http.StatusInternalServerError){ return }
	//сохранение в базу
	query := "INSERT INTO users (name, password, email) VALUES (?, ?, ?)"
	_, err = database.DB.Exec(query, user.Name, string(hashedPassword), user.Email)
	if !utils.CheckError(w, err, "Ошибка регистрации пользователя", http.StatusInternalServerError){ return }
	//отправка ответа
	utils.ReturnResponse(w, map[string]string{"message":"регистрация прошла успешно!"}, http.StatusCreated)
}

//login
func LoginHandler(w http.ResponseWriter, r *http.Request){
	//проверка метода запроса
	if !utils.CheckMethod(r, w, http.MethodPost){ return }
	//декодирование данных
	var user models.User
	if !utils.DecodeData(r, w, &user){ return }
	//выполнение запроса - поиск пользователя в базе и получение хэшированного пароля
	var hashedPassword string
	query := "SELECT password FROM users WHERE email = ?"
	err := database.DB.QueryRow(query, user.Email).Scan(&hashedPassword)
	if err != nil {
    if err == sql.ErrNoRows {
        utils.CheckError(w, err, "Неверный email или пароль", http.StatusUnauthorized)
    } else {
        utils.CheckError(w, err, "Ошибка сервера", http.StatusInternalServerError)
    }
    return
	}
	//проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if !utils.CheckError(w, err, "Неверный email или пароль", http.StatusUnauthorized){ return }
	//генерация токена
	tokenString, err := GenerateToken(user.Email)
	if !utils.CheckError(w, err, "Ошибка генерации токена", http.StatusInternalServerError){ return }
	//отправка ответа - сообщение и токен
	utils.ReturnResponse(w, map[string]string{"message":"авторизация прошла успешно!", "token": tokenString}, http.StatusOK)
}

//login user

