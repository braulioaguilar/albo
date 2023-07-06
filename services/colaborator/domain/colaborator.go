package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Colaborator struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Editor    []string           `bson:"editors"`
	Writer    []string           `bson:"writers"`
	Colorist  []string           `bson:"colorists"`
}
