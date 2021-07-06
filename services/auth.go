package services

import (
	"context"
	"fmt"
	"github.com/its-dastan/go-blog/db"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	dbs  = "go-blog"
	coll = "users"
)

func Register(user map[string]interface{}) map[string]interface{} {
	client := db.Connect()
	defer db.DisConnect(context.Background(), client)
	//hashedPassword, err := helper.EncryptPassword(user["password"].(string))
	//if err!= nil{
	//	panic(err.Error())
	//}
	//fmt.Println(string(hashedPassword))
	collection := client.Database(dbs).Collection(coll)
	var result bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"email": user["email"]}).Decode(&result); err != nil {
		userT, _ := bson.Marshal(user)
		_, err := collection.InsertOne(context.Background(), userT)
		if err != nil {
			return map[string]interface{}{
				"Status": false,
				"Msg":    "Account couldn't be created. Please try again!",
			}
		}

		if err := collection.FindOne(context.TODO(), bson.M{"email": user["email"]}).Decode(&result); err != nil {
			return map[string]interface{}{
				"Status": false,
				"Msg":    "Error while getting the user",
			}
		}
		return map[string]interface{}{
			"Status": true,
			"Msg":    "Successfully Registered",
			"Data":   result,
		}
	} else {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Email Already exists",
		}
	}
}

func Login(user map[string]interface{}) map[string]interface{} {
	client := db.Connect()
	defer db.DisConnect(context.Background(), client)
	collection := client.Database(dbs).Collection(coll)
	var result bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"email": user["email"]}).Decode(&result); err != nil {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Invalid Email id",
		}
	}
	fmt.Println(result["password"], user["password"])
	if user["password"] != result["password"] {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Wrong Password",
		}
	}
	return map[string]interface{}{
		"Status": true,
		"Msg":    "Successfully logged in",
		"Data":   result,
	}
}

func GetUsers() map[string]interface{} {
	client := db.Connect()
	defer db.DisConnect(context.Background(), client)
	collection := client.Database(dbs).Collection(coll)
	curr, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Internal Server Error",
		}
	}
	var result []bson.M
	if err := curr.All(context.Background(), &result); err != nil {
		return map[string]interface{}{
			"Status": false,
			"Msg":    "Internal Server Error",
		}
	}
	return map[string]interface{}{
		"Status": true,
		"Msg":    "Successfully Registered",
		"Data":   result,
	}

	//fmt.Printf("%T\n", result)
	//fmt.Println(result)
}
