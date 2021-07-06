package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/services"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerData map[string]interface{}
	var result map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&registerData)

	if err != nil || registerData["firstName"] == "" || registerData["lastName"] == "" || registerData["password"] == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please Enter all the data correctly"})
		return
	}
	err = services.Register(registerData, &result)
	if err!= nil{
		panic(err)
	}
	fmt.Println(result)
	//if response["Status"] == false {
	//	helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: response["Msg"]})
	//} else {
	//	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: response["Msg"], Data: response["Data"]})
	//}
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Hello")
//	var loginData map[string]interface{}
//	err1 := json.NewDecoder(r.Body).Decode(&loginData)
//	if err1 == nil{
//		fmt.Println(loginData)
//	}
//	if loginData["email"] == "" || loginData["password"] == "" {
//		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please Enter all the data correctly"})
//		return
//	}
//	response := services.Login(loginData)
//	if response["Status"] == false {
//		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: response["Msg"]})
//	} else{
//		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "successful", Data: response["Data"]})
//	}
//}
//
//func GetUsers(w http.ResponseWriter, r *http.Request) {
//	response := services.GetUsers()
//	if response["Status"] == false {
//		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: response["Msg"]})
//	}
//	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: response["Msg"], Data: response["Data"]})
//
//}
