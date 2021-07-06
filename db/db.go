package db

import (
	"github.com/globalsign/mgo"
	"log"
)

const (
	host     = "mongodb://localhost:27017"
	database = "go-blog"
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Database: database,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln(err)
	}
	globalS = s
}

func Connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB("").C(collection)
	return s, c
}
