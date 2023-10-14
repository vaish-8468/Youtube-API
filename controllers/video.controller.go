package controllers

import (
	"FamPay/models"
	"FamPay/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//interacts with user service

type VideoController struct{
	VideoServices services.VideoService //reference of videoservice
}

func NewVideo(videoservice services.VideoService) VideoController{
	return VideoController{
		VideoServices: videoservice,

	}
}

//from controller we'll define routes , hence we'll call those methods defines in models
//gin.context holds information about the request that we're gonna send and will get a json object response
func (uc *VideoController) CreateList(ctx *gin.Context) {
	var video models.Video
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&video); err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	//we'll check for the error while creating new user
	err := uc.VideoServices.CreateList(&video) //passing the address of the user object
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

//route functions or handlers
func (uc *VideoController) GetList(ctx *gin.Context){
	title :=ctx.Param("title") //fetching the name from ctx object
	user, err := uc.VideoServices.GetList(&title)
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK,user) //return the user object
}


func (uc *VideoController) GetAll(ctx *gin.Context){
	videos, err := uc.VideoServices.GetAll()
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK,videos) //return the user object
}


func (uc *VideoController) UpdateList(ctx *gin.Context){
	var video models.Video
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&video); err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	err := uc.VideoServices.UpdateList(&video)
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK,video) //return the user object
}



func (uc *VideoController) DeleteList(ctx *gin.Context){
	title := ctx.Param("title")
	err := uc.VideoServices.DeleteList(&title)
	if err!=nil{
			//error while saving in the mongodb
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
		}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


//route function to group all the routes
func (uc *VideoController) RegisterVideoRoutes(rg *gin.RouterGroup){
	videoRoute :=rg.Group("/video")  //base path
	videoRoute.POST("/create", uc.CreateList)
	videoRoute.GET("/get/:title", uc.GetList)
	videoRoute.GET("/getAll", uc.GetAll)
	videoRoute.PATCH("/update", uc.UpdateList)
	videoRoute.DELETE("/delete:title", uc.DeleteList)
}