package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

var client = &mongo.Client{}
var mainDB = ""

// InitDB This function will connect to the db provided on the params.
func InitDB(dbURL, dbPort, dbUser, dbPass, dbAuth, dbMain string) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbURL + ":" + dbPort).SetAuth(options.Credential{
		Username:   dbUser,
		Password:   dbPass,
		AuthSource: dbAuth, // db name
	})
	mainDB = dbMain
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Mongo Connection Error: " + err.Error())
	}

}

// GetClient  get the client provided by the configuration. singleton for connection pool.
func GetClient() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	return nil, errors.New("cannot-connect-db")

}

// GetDatabase  get the database obj provided by the configuration. singleton for connection pool.
func GetDatabase() (*mongo.Database, error) {
	if client != nil {
		return client.Database(mainDB), nil
	}

	return nil, errors.New("cannot-get-db")

}

// GetCollection  get the coll obj provided by the configuration. singleton for connection pool.
func GetCollection(collName string) (*mongo.Collection, error) {
	if client != nil {
		return client.Database(mainDB).Collection(collName), nil
	}

	return nil, errors.New("cannot-get-collection")

}
