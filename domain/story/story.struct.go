package story

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type File struct {
	Url  string `json:"text" bson:"text"`
	Size int    `json:"size" bson:"int"`
}

type Story struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Caption   string             `json:"caption" bson:"caption"`
	Text      string             `json:"text" bson:"text,omitempty"`
	File      File               `json:"file" bson:"file,omitempty"`
	UserId    string             `json:"userId" bson:"userId"`
	CreatedAt *time.Time         `json:"createdAt" bson:"createdAt"`
}
