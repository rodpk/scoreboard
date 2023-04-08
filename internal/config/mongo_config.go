package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMongoConfig(host string, port int, username string, password string) *MongoConfig {
	return &MongoConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (mc *MongoConfig) Connect() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", mc.Username, mc.Password, mc.Host, mc.Port)
	clientOptions := options.Client().ApplyURI(uri)

	// connect to mongo
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
