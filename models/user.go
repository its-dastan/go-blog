package models

type User struct {
	FirstName string      `bson:"firstName" json:"firstName"`
	LastName  string      `bson:"lastName" json:"lastName"`
	Email     string      `bson:"email" json:"email"`
	Password  string      `bson:"password" json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}
