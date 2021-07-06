package helper

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func EncryptPassword(password string)([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), 15)
}