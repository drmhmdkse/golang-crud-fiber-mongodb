package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}
