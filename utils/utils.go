package utils

import (
	"encoding/json"
	"net/http"
)

//check error
func CheckError(w http.ResponseWriter, err error, text string, status int)bool{
	if err != nil {
		http.Error(w, text, status)
		return false
	}
	return true
}

//check methods
func CheckMethod(r *http.Request, w http.ResponseWriter, methodServer string)bool{
	if r.Method != methodServer{
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

//decode data
func DecodeData(r *http.Request, w http.ResponseWriter, dst interface{})bool{
	err:= json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return false
	}
	return true
}

//return response
func ReturnResponse(w http.ResponseWriter, message map[string]string, status int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}
