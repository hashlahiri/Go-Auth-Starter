package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	
	"go-auth-starter.app/base/config"
	"go-auth-starter.app/base/domain"
	"go-auth-starter.app/base/service"
	
)

func closeDBConnectionAfterShutdown(client *mongo.Client) {
	
	// Ensure the MongoDB client disconnects when the server shuts down
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()
}

func connectToDB() *mongo.Client {

	// Initialize MongoDB connection
	client, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	// Set the event collection in the domain package
	domain.SetUserCollection(client)
	return client
}

func main() {

	mongoClient := connectToDB() // Connect to DB

	server := gin.Default() // gives http server + middleware

	service.ApiEndpoints(server) // all api-endpoints here

	server.Run(":9000") // to start the server
	closeDBConnectionAfterShutdown(mongoClient) // close DB connection after shutdown
}