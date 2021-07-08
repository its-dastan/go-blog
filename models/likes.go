package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Likes struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	BlogId  bson.ObjectId `json:"blogId,omitempty" bson:"blogId,omitempty"`
	LikedBy bson.ObjectId `json:"likedBy,omitempty" bson:"likedBy,omitempty"`
	LikedAt time.Time     `json:"likedAt,omitempty" bson:"likedAt,omitempty"`
}
