package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)
var DB *mongo.Database

func DbConnect () *mongo.Client{
	mongoURL :=  os.Getenv("DB_URL");
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(mongoURL))

	if err != nil{
		log.Fatal("Error while connect to db", err)
	}
	DB = mongoClient.Database("odin")
	log.Println("Connected to DB");

	err = mongoClient.Ping(context.Background(),readpref.Primary())
	if err != nil{
		log.Fatal("ping failed", err)
	}
	log.Println("ping success");

	return mongoClient;
}