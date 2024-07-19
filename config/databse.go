// config/database.go
package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client    *mongo.Client
	UserColl  *mongo.Collection
	OrderColl *mongo.Collection
	Ctx       context.Context
)

func ConnectDatabase() {
	Ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(Ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = Client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	UserColl = Client.Database("userdb").Collection("users")
	OrderColl = Client.Database("orderdb").Collection("orders")
}

func DisconnectDatabase() {
	if err := Client.Disconnect(Ctx); err != nil {
		log.Fatal("error while disconnecting from mongo", err)
	}
}
