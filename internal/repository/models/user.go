package models

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Admin    bool   `bson:"admin"`
}
