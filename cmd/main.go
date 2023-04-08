package main

import (
	"context"
	"log"

	"github.com/rodpk/scoreboard/internal/config"
	"github.com/rodpk/scoreboard/internal/repository"
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

	// todo: inject depen
	repository.NewScoreboardRepository(client)




}
