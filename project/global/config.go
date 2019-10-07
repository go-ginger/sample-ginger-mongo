package global

import (
	gh "github.com/go-ginger/helpers"
	gm "github.com/go-ginger/models"
	"github.com/go-ginger/mongo"
)

type config struct {
	gm.IConfig

	Port string

	// mongo config
	MongoConnectionString string
	MongoDatabaseName     string
	//
}

var Config *config

func initializeConfig() {
	Config = &config{
		Port:                  gh.GetEnv("Port", "8080"),
		MongoConnectionString: gh.GetEnv("MongoConnectionString", "mongodb://localhost:27017"),
		MongoDatabaseName:     gh.GetEnv("MongoDatabaseName", "test"),
	}
	mongo.InitializeConfig(Config)
}
