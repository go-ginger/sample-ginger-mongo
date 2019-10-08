package models

import (
	"github.com/go-ginger/mongo"
	"time"
)

type Movie struct {
	mongo.BaseModel `bson:"inline"`

	Title       string    `bson:"title,omitempty" json:"title,omitempty"`
	Rate        float64   `bson:"rate,omitempty" json:"rate,omitempty"`
	ReleaseDate time.Time `bson:"release_date,omitempty" json:"release_date,omitempty"`
	Synopsis    string    `bson:"synopsis,omitempty" json:"synopsis,omitempty"`
}
