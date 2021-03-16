package main

import (
	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/app"
	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
