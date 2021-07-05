package services

import (
	"context"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	dbs  = "go-blog"
	coll = "users"
)

func Register(user models.User) map[string]interface{} {
	client := db.Connect()
	defer db.DisConnect(context.Background(), client)
	collection := client.Database(dbs).Collection(coll)
	var result bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result); err != nil {
		userT, _ := bson.Marshal(user)
		_, err := collection.InsertOne(context.Background(), userT)
		if err != nil {
			return map[string]interface{}{
				"Status": false,
				"Msg":    "Account couldn't be created. Please try again!",
			}
		}
		var result1 bson.M

		if err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result1); err != nil {
			return map[string]interface{}{
				"Status": false,
				"Msg":    "Error while getting the user",
			}
		}
		return map[string]interface{}{
			"Status": true,
			"Msg":    "Successfully Registered",
			"Data":   result1,
		}
	} else {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Email Already exists",
		}
	}
}
