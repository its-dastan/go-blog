package controllers

import (
	"encoding/json"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/services"
	"net/http"
)


func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.FirstName == "" || user.LastName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please Enter all the data correctly"})
		return
	}
	response := services.Register(user)
	if response["Status"] == false {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: response["Msg"]})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: response["Msg"], Data: response["Data"]})
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	response := services.GetUsers()
	if response["Status"] == false{
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: response["Msg"]})
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK,Msg: response["Msg"], Data: response["Data"]})
}
