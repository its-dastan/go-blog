package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/service"
	"net/http"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Password  string        `json:"password" bson:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerData map[string]interface{}
	var result User
	err := json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil || registerData["firstName"] == "" || registerData["lastName"] == "" || registerData["email"] == "" || registerData["password"] == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
	}
	err = service.Register(registerData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
	} else {
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData map[string]interface{}
	var result User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil || loginData["email"] == "" || loginData["password"] == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please type all field"})
	}
	err = service.Login(loginData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
	} else {
		fmt.Printf("%v\n", result.ID.Hex())
		helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
	}
}
