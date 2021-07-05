package services

import (
	"context"
	"fmt"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

const (
	dbs  = "go-blog"
	coll = "users"
)

func Register(user models.User) {
	fmt.Println(user)
	client := db.Connect()
	defer db.DisConnect(context.Background(), client)
	collection := client.Database(dbs).Collection(coll)
	curr, err := collection.Find(context.Background(), bson.M{"email": user.Email})
	if err != nil {
		log.Fatal(err)
	}
	defer curr.Close(context.Background())
	var result []bson.M
	if err := curr.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	if len(result) > 0 {
		fmt.Println("Hello")

		fmt.Println("Email already exists")
	} else {
		userT, _ := bson.Marshal(user)
		res, err := collection.InsertOne(context.Background(), userT)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	}

}
