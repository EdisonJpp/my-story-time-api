package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Password    string             `json:"password" bson:"password"`
}
