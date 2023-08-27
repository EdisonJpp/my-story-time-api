package user

import "go.mongodb.org/mongo-driver/bson"

type UserRepository interface {
	GetOneBy(filters *bson.M) (*User, error)
}
