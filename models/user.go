package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string `form:"username" json:"username" bson:"username" validate:"required,min=3"`
	Email    string `form:"email" json:"email" bson:"email" validate:"required,email"`
	Password string `form:"password" json:"password" bson:"password" validate:"required,min=6"`
	CreatedAt string `form:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt string `form:"updated_at" json:"updated_at" bson:"updated_at"`
}