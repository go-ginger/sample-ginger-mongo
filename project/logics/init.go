package logics

import (
	l "github.com/go-ginger/logic"
	"github.com/go-ginger/sample-ginger-mongo/project/db"
)

var MovieLogicHandler movieLogic

func init()  {
	MovieLogicHandler = movieLogic{IBaseLogicHandler: &l.BaseLogicHandler{}}

	MovieLogicHandler.Init(&MovieLogicHandler, &db.MovieDbHandler)
}
