package services

import (
	"github.com/globalsign/mgo/bson"
	"github.com/its-dastan/go-blog/db"
)

const (
	dbs = "go-blog"
	col = "users"
)

func Register(user map[string]interface{}, result interface{}) error {
	s, c := db.Connect(col)
	defer s.Close()
	err:= c.Insert(user)
	if err != nil{
		panic(err)
	}
	return c.Find(bson.M{"email": user["email"]}).One(result)
}
