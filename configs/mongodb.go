package configs

import (
	// "FamPay/controllers"
	// "FamPay/services"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

	var (
		// vs          services.VideoService
		// vc          controllers.VideoController
		ctx         context.Context
		// videoc      *mongo.Collection
		mongoclient *mongo.Client
		err         error
	)
	


func Database()(*mongo.Client,error){
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoconn := options.Client().ApplyURI("mongodb+srv://2202vartikavsh:vart987654321@cluster0.vu5eoiq.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// mongoconn := options.Client().ApplyURI(configs.EnvMongoURI()).SetServerAPIOptions(serverAPI)
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

	//logic for fetching videos
	// videoc = mongoclient.Database("userdb").Collection("videos")

	return mongoclient,err

}