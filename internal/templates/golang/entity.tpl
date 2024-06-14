package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type {{.className}} struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (a *{{.className}}) Prepare() {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}