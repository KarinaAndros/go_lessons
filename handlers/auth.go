package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
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
	if r.Method != http.MethodPost{
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	//декодирование JSON из тела запроса
	var user models.User
	err:= json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	//хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Ошибка хэширования пароля", http.StatusInternalServerError)
		return
	}
	//сохранение в базу
	query := "INSERT INTO users (name, password, email) VALUES (?, ?, ?)"
	_, err = database.DB.Exec(query, user.Name, string(hashedPassword), user.Email)
	if err != nil {
		log.Println("Ошибка БД", err)
		http.Error(w, "Ошибка регистрации пользователя", http.StatusInternalServerError)
		return
	}
	//отправка ответа
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"Пользователь зарегистрирован"})
}

//login
func LoginHandler(w http.ResponseWriter, r *http.Request){
	//проверка метода запроса
	if r.Method != http.MethodPost{
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
		return
	}
	//декодирование данных
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	//выполнение запроса - поиск пользователя в базе и получение хэшированного пароля
	var hashedPassword string
	query := "SELECT password FROM users WHERE email = ?"
	err = database.DB.QueryRow(query, user.Email).Scan(&hashedPassword)
	if err != nil{
		if err == sql.ErrNoRows{
			http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	//проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil{
		http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
		return
	}
	//генерация токена
	tokenString, err := GenerateToken(user.Email)
	if err != nil{
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}
	//отправка ответа - сообщение и токен
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"авторизация прошла успешно!", "token": tokenString})
}
