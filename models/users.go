package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	ID         bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName  string          `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName   string          `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email      string          `json:"email,omitempty" bson:"email,omitempty"`
	Password   string          `json:"password,omitempty" bson:"password,omitempty"`
	LikedBlogs []bson.ObjectId `json:"likedBlogs,omitempty" bson:"likedBlogs,omitempty"`
	CreatedAt  time.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
