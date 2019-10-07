package v1

import (
	g "github.com/go-ginger/ginger"
	"github.com/go-ginger/sample-ginger-mongo/project/db"
	"github.com/go-ginger/sample-ginger-mongo/project/logics"
)

var movies = new(MoviesController)
var movie = new(MovieController)

func init() {
	movies.Init(movies, &logics.MovieLogicHandler, &db.MovieDbHandler)
	movie.Init(movie, &logics.MovieLogicHandler, &db.MovieDbHandler)
}

func RegisterRoutes(router *g.RouterGroup) {
	movies.AddRoute("Post")
	movies.AddRoute("Get")

	movie.AddRoute("Get")
	movie.AddRoute("Put")
	movie.AddRoute("Delete")

	movies.RegisterRoutes(movies, "/movies", router)
	movie.RegisterRoutes(movie, "/movies/:id", router)
}
