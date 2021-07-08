package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Comments struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Comment     bson.ObjectId `json:"comment,omitempty" bson:"comment,omitempty"`
	BlogId      bson.ObjectId `json:"blogId,omitempty" bson:"blogId,omitempty"`
	CommentedBy bson.ObjectId `json:"commentedBy,omitempty" bson:"commentedBy,omitempty"`
	CommentedAt time.Time     `json:"commentedAt,omitempty" bson:"commentedAt,omitempty"`
}
