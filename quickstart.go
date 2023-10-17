//initialize everything and start the gin server

package main

import (
	// "FamPay/configs"
	"FamPay/controllers"
	"FamPay/models"
	// "os"
	"time"

	"FamPay/services"
	"context"

	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/googleapi/transport"
	// "google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	server      *gin.Engine //server for the gin framework
	us          services.UserService
	uc          controllers.UserController
	vs          services.VideoService
	vc          controllers.VideoController
	ctx         context.Context
	userc       *mongo.Collection
	videoc      *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

var (
	query      = flag.String("query", "Lenovo", "Search term")
	maxResults = flag.Int64("max-results", 10, "Max YouTube results")
)


const developerKey = "AIzaSyCVyAuztulPhUTKsskORtZI6RmsLl5TWlk"

// var developerKey=os.Getenv("APIkey")
//  var MongoURI=os.Getenv("MongoURI")
//  var username=os.Getenv("USERNAME")
//  var password=os.Getenv("PASSWORD")

// we'll initialize them in the init function
func init() {
	flag.Parse()
	ctx = context.TODO() //will create a single context object with no cancellation thing inside

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
		// Transport: &transport.APIKey{Key: configs.APIkey()},
	}
	// service, err := youtube.NewService(ctx, option.WithAPIKey(configs.APIkey()))
	service, err := youtube.New(client)
	//  // Create a new HTTP client.
	//  httpClient := &http.Client{}

	//  // Initialize the YouTube service with the API key.
	//  service, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient), option.WithAPIKey(configs.APIkey()))
	 
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		fmt.Println("Error encountered")
	}

	

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

	videoc = mongoclient.Database("userdb").Collection("videos")

	// Initialize a slice to store the video information.
	var list []models.Video

	// Iterate through each item in the response and store the relevant data.
	for _, video := range response.Items {
		// Create a Video instance for each video in the response
		publishedAt,err:=time.Parse(time.RFC3339, video.Snippet.PublishedAt)
		if err != nil {
			log.Println(err)
		}
		videos := models.Video{
			Kind: video.Kind,
			Etag: video.Etag,
			Id: models.Id{
				IDKind:     video.Id.Kind,
				VideoId:    video.Id.VideoId,
				ChannelId:  video.Id.ChannelId,
				PlaylistId: video.Id.PlaylistId,
			},
			Snippet: models.Snippet{
				PublishedAt:      publishedAt,
				SnippetChannelId: video.Snippet.ChannelId,
				Title:            video.Snippet.Title,
				Description:      video.Snippet.Description,
				// ThumbnailUrl: video.Snippet.Thumbnails["default"], // Access the default thumbnail
				ChannelTitle: video.Snippet.ChannelTitle,
				// LiveBroadcastContent: video.Snippet.LiveBroadcastContent,
			},
		}
		list = append(list, videos)
		// Insert the video document into MongoDB
		_, err = videoc.InsertOne(ctx, videos)
		if err != nil {
			log.Println(err)
		}
		//creating text indexing
		model := mongo.IndexModel{Keys: bson.D{{Key: "snippet.title", Value: "text"},{Key: "snippet.description", Value: "text"}}}
		_, err = videoc.Indexes().CreateOne(context.TODO(), model)
		if err != nil {
			panic(err)
		}

	}
	//initializing video collection
	vs = services.NewVideoService(videoc, ctx)
	vc = controllers.NewVideo(vs)
	

	//dropping the existing indices
		// _, err = videoc.Indexes().DropOne(context.TODO(), "snippet.title_text_snippet.description_text")
	// _, err = videoc.Indexes().DropOne(context.TODO(), "Snippet.title_text_Snippet.description_text")


	//initialize usercollection
	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)

	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)

	userBasePath := server.Group("/v1")
	videoBasePath :=server.Group("/v2")


	uc.RegisterUserRoutes(userBasePath)
	vc.RegisterVideoRoutes(videoBasePath)

	log.Fatal(server.Run(":9090"))

}
