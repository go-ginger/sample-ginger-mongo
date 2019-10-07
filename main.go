package main

import (
	"github.com/go-ginger/sample-ginger-mongo/project/controllers"
	"github.com/go-ginger/sample-ginger-mongo/project/db"
	"github.com/go-ginger/sample-ginger-mongo/project/global"
	"log"
)

func main() {
	global.Initialize()
	db.Initialize()
	engine := controllers.RegisterRoutes()
	err := engine.Run()
	log.Fatal(err)
}
