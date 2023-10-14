// package main

// import (
//         // "flag"
//         "fmt"
//         // "log"
//         "net/http"
// 		// "go.mongodb.org/mongo-driver/bson/primitive"
// 		"io/ioutil"

//         // "google.golang.org/api/googleapi/transport"
//         // "google.golang.org/api/youtube/v3"
// )

// func main() {
//     url := "https://www.googleapis.com/youtube/v3/videos?id=7lCDEYXw3mM&key=AIzaSyCVyAuztulPhUTKsskORtZI6RmsLl5TWlk&part=snippet,contentDetails,statistics,status"
//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         fmt.Print(err.Error())
//     }
//     res, err := http.DefaultClient.Do(req)
//     if err != nil {
//         fmt.Print(err.Error())
//     }
//     defer res.Body.Close()
//     body, readErr := ioutil.ReadAll(res.Body)
//     if readErr != nil {
//         fmt.Print(err.Error())
//     }
//     fmt.Println(string(body))
// }

//initialize everything and start the gin server

package main

import (
	"FamPay/controllers"
	"FamPay/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	server          *gin.Engine  //server for the gin framework 
	us     			services.UserService
	uc  			controllers.UserController
	videoservice    services.VideoService
	videocontroller controllers.VideoController
	ctx             context.Context
	userc  			*mongo.Collection
	videocollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error

)

//we'll initialize them in the init function
func init(){
	ctx = context.TODO() //will create a single context object with no cancellation thing inside

	// mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoconn := options.Client().ApplyURI("mongodb+srv://2202vartikavsh:vart987654321@cluster0.vu5eoiq.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	mongoclient, err := mongo.Connect(context.TODO(), mongoconn)
	if err != nil {
	  panic(err)
	}
	// defer func() {
	//   if err = mongoclient.Disconnect(context.TODO()); err != nil {
	// 	panic(err)
	//   }
	// }()
	// Send a ping to confirm a successful connection
	if err := mongoclient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You're successfully connected to MongoDB!")
  
  
	// mongoclient, err = mongo.Connect(ctx, mongoconn)
	// if err!=nil{
	// 	//close application by passing error in the console
	// 	log.Fatal(err)
	// }

	// //if there is error while pinging the mongodb at first then
	// err= mongoclient.Ping(ctx, readpref.Primary()) //primary database
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println("mongo connection has been established!")

	//initialize usercollection
	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc , ctx)
	uc = controllers.New(us)
	server=gin.Default()

}

func main(){
defer mongoclient.Disconnect(ctx)

basePath:= server.Group("/v1")

uc.RegisterUserRoutes(basePath)

log.Fatal(server.Run(":9090"))

}
