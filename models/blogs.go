package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Blog struct {
	ID        bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
	Caption   string          `json:"caption,omitempty" bson:"caption,omitempty"`
	Image     string          `json:"image,omitempty" bson:"image,omitempty"`
	Video     string          `json:"video,omitempty" bson:"video,omitempty"`
	Likes     []bson.ObjectId `json:"likes,omitempty" bson:"likes,omitempty"`
	Comments  []bson.ObjectId `json:"comments,omitempty" bson:"comments,omitempty"`
	CreatedAt time.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	PostedBy  bson.ObjectId   `json:"postedBy,omitempty" bson:"postedBy,omitempty"`
}
