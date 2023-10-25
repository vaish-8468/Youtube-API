//initialize everything and start the gin server

package main

import (
	"Youtube_RestAPI/configs"
	"Youtube_RestAPI/controllers"
	"Youtube_RestAPI/services"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine //server for the gin framework
	vs     services.VideoService
	vc     controllers.VideoController
	ctx    context.Context
	videoc *mongo.Collection
)

// we'll initialize them in the init function
func init() {

	ctx = context.TODO()
	//will create a single context object with no cancellation thing inside

	mongoclient, err := configs.Database()
	if err != nil {
		panic(err)
	}

	videoc = mongoclient.Database("userdb").Collection("videos")

	//initializing video collection
	vs = services.NewVideoService(videoc, ctx)
	vc = controllers.NewVideo(vs)

	server = gin.Default()

}

func main() {

	videoBasePath := server.Group("/v2")
	vc.RegisterVideoRoutes(videoBasePath)
	log.Fatal(server.Run(":9090"))

}
