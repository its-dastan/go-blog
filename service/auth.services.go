package service

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
)

const (
	usersCollection = "users"
)

func Register(user *models.User, result interface{}) error {
	s, c := db.Connect(usersCollection)
	defer s.Close()
	count, err := c.Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email already exists")
	}
	userCopy := &models.User{
		Email:    user.Email,
		Password: user.Password,
	}
	hashedPassword, err := helper.EncryptPassword(user.Password)
	if err != nil {
		return errors.New("please choose a different password")
	}
	user.Password = string(hashedPassword)
	err = c.Insert(user)
	if err!=nil{
		return errors.New("internal error! please try again later")
	}
	return Login(userCopy, result)
}

func Login(user *models.User, result interface{}) error {
	s, c := db.Connect(usersCollection)
	defer s.Close()
	var result1 *models.User
	_ = c.Find(bson.M{"email": user.Email}).One(&result1)
	if result1 == nil {
		return errors.New("invalid email id")
	}
	err := helper.ComparePassword(result1.Password, user.Password)
	if err != nil {
		return errors.New("wrong password")
	}
	return c.Find(bson.M{"email": user.Email}).One(result)
}
