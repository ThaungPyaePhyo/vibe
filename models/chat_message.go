package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatMessage struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	SenderID  primitive.ObjectID `json:"sender_id" bson:"sender_id"`
	ReceiverID primitive.ObjectID `json:"receiver_id" bson:"receiver_id"`
	Message  string             `json:"message" bson:"message"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
	IsRead    bool               `json:"is_read" bson:"is_read"`
}