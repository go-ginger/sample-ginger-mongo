package v1

import (
	"github.com/go-ginger/sample-ginger-mongo/project/models"
)

type MoviesController struct {
	models.BaseItemsController
}

type MovieController struct {
	models.BaseItemController
}
