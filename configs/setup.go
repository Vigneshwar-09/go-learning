package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
    defer cancel()
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("development").Collection(collectionName)
}
