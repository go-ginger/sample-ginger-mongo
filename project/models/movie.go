package models

import (
	"github.com/go-ginger/mongo"
	"time"
)

type Movie struct {
	mongo.BaseModel

	Title       string    `bson:"title" json:"title,omitempty"`
	Rate        float64   `bson:"rate" json:"rate,omitempty"`
	ReleaseDate time.Time `bson:"release_date" json:"release_date,omitempty"`
	Synopsis    string    `bson:"synopsis" json:"synopsis,omitempty"`
}
