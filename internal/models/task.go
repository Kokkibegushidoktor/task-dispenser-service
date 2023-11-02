package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}

type TaskLevel struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TaskId        primitive.ObjectID `json:"taskId" bson:"taskId"`
	Title         string             `json:"title" bson:"title"`
	VarQuestCount int                `json:"varQuestCount" bson:"varQuestCount"`
	Questions     []LevelQuestion    `json:"questions" bson:"questions"`
}

type LevelQuestion struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
}
