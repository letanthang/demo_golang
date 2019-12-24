package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func main() {
	connect()
	listenChangeStream(Client.Database("golang11").Collection("student"))
}

func listenChangeStream(collection *mongo.Collection) {
	ctx := context.Background()
	stream, err := collection.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		log.Println(err)
		return
	}
	defer stream.Close(ctx)

	for stream.Next(ctx) {
		var elem bson.M

		if err := stream.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Println(elem)
	}
}

func connect() {
	uri := "mongodb+srv://mongoadmin:secret1234@cluster0-xxyrd.gcp.mongodb.net"
	log.Println("Try to connect mongo:", uri)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(4)
	clientOptions.SetReadPreference(readpref.Nearest())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}
	Client = client
}
