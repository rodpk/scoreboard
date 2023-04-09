package main

import (
	"context"
	"log"

	"github.com/rodpk/scoreboard/internal/config"
	"github.com/rodpk/scoreboard/internal/handler"
	"github.com/rodpk/scoreboard/internal/repository"
	"github.com/rodpk/scoreboard/internal/server"
	"github.com/rodpk/scoreboard/internal/routes"
)

func main() {

	// create new mongodb
	mongoConfig := config.NewMongoConfig("localhost", 27017, "root", "example")

	// connect
	client, err := mongoConfig.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	router := server.CreateServer()

	scoreboardRepository := repository.NewScoreboardRepository(client)
	scoreboardHandler := handler.NewScoreboardHandler(scoreboardRepository)
	routes.InitializeRoutes(router, scoreboardHandler)
	
	server.StartServer(router)
}
