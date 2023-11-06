package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Admin    bool               `json:"admin" bson:"admin"`
	Roles    []string           `json:"roles,omitempty" bson:"roles,omitempty"`
}
