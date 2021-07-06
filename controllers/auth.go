package controllers

import (
	"encoding/json"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/service"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerData map[string]interface{}
	var result map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil || registerData["firstName"] == "" || registerData["lastName"] == "" || registerData["email"] == "" || registerData["password"] == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
	}
	err = service.Register(registerData, &result)
	if err!=nil{
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
	} else{
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	var loginData map[string]interface{}
	var result map[string]interface{}
	err:= json.NewDecoder(r.Body).Decode(&loginData)
	if err!=nil || loginData["email"]==""|| loginData["password"]==""{
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
	}
	err = service.Login(loginData, &result)
	if err!=nil{
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
	} else{
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
	}
}
