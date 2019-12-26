package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

var client = &mongo.Client{}

// InitDB This function will connect to the db provided on the params.
func InitDB(dbURL, dbPort, dbUser, dbPass, dbAuth string) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbURL + ":" + dbPort).SetAuth(options.Credential{
		Username:   dbUser,
		Password:   dbPass,
		AuthSource: dbAuth, // db name
	})
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Mongo Connection Error: " + err.Error())
	}

}

// GetConnection  get the client provided by the configuration. singleton for connection pool.
func GetConnection() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	return nil, errors.New("Cannot connect to DB")

}
