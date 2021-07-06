package service

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/helper"
)

const (
	coll = "users"
)

func Register(user map[string]interface{}, result interface{}) error {
	s, c := db.Connect(coll)
	defer s.Close()
	count, err := c.Find(bson.M{"email": user["email"]}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("email already exists")
		return errors.New("email already exists")
	}
	hashedPassword, err := helper.EncryptPassword(user["password"].(string))
	if err != nil {
		return errors.New("please choose a different password")
	}
	user["password"] = string(hashedPassword)
	_ = c.Insert(user)
	return c.Find(bson.M{"email": user["email"]}).One(result)
}

func Login(user map[string]interface{}, result interface{}) error {
	s, c := db.Connect(coll)
	defer s.Close()
	var result1 map[string]interface{}
	_ = c.Find(bson.M{"email": user["email"]}).One(&result1)
	if result1 == nil {
		return errors.New("invalid email id")
	}
	err := helper.ComparePassword(result1["password"].(string), user["password"].(string))
	if err != nil {
		return errors.New("wrong password")
	}
	return c.Find(bson.M{"email": user["email"]}).One(result)
}
