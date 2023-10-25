package configs

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

		

func Database()(*mongo.Client,error){
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")	
	connectionString := "mongodb+srv://" + username + ":" + password + "@cluster0.vu5eoiq.mongodb.net/?retryWrites=true&w=majority"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoconn := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	mongoclient, err := mongo.Connect(context.TODO(), mongoconn)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := mongoclient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You're successfully connected to MongoDB!")
	return mongoclient,err

}