package controllers

import (
	"encoding/json"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerData *models.User
	var result *models.User
	err := json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil || registerData.FirstName == "" || registerData.LastName == "" || registerData.Email == "" || registerData.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
		return
	}
	err = service.Register(registerData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})

}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData *models.User
	var result *models.User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil || loginData.Email == "" || loginData.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
		return
	}
	err = service.Login(loginData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
}
