package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/benchungus/api/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//connects the user to the db
func ConnDB() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	//create new client with uri
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	//connecting to client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//ping client to verify connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	utils.Logger.Info("connected to db")
	return client
}
