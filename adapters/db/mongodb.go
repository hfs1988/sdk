package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	URI string
}

func GetMongoDBInstance(uri string) *mongoDB {
	return &mongoDB{
		URI: uri,
	}
}

func (m *mongoDB) Connect() error {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(m.URI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	defer func() error {
		if err = client.Disconnect(context.TODO()); err != nil {
			return err
		}
		return nil
	}()

	fmt.Println("You successfully connected to MongoDB!")
	return nil
}
