package db

import "github.com/go-ginger/sample-ginger-mongo/project/models"

var MovieDbHandler = movieMongoHandler{}

func Initialize() {
	MovieDbHandler.Initialize(models.Movie{})
}
