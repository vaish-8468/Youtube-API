package controllers

import (
	"Youtube_RestAPI/models"
	"Youtube_RestAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//interacts with user service

type VideoController struct {
	VideoServices services.VideoService //reference of videoservice
}

func NewVideo(videoservice services.VideoService) VideoController {
	return VideoController{
		VideoServices: videoservice,
	}
}

// from controller we'll define routes , hence we'll call those methods defines in models
// gin.context holds information about the request that we're gonna send and will get a json object response
func (vc *VideoController) CreateList(ctx *gin.Context) {
	var video models.Video
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	//we'll check for the error while creating new user
	err := vc.VideoServices.CreateList(&video) //passing the address of the user object
	if err != nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// route functions or handlers
func (vc *VideoController) GetListByQuery(ctx *gin.Context) {
	query := ctx.Param("query") //fetching the name from ctx object
	page := ctx.Param("page")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	video, err := vc.VideoServices.GetList(&query, &page, &pageSize)

	if err != nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK, video) //return the user object
}

func (vc *VideoController) GetAll(ctx *gin.Context) {
	videos, err := vc.VideoServices.GetAll()
	if err != nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK, videos) //return the user object
}

func (vc *VideoController) UpdateList(ctx *gin.Context) {
	var video models.Video
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	err := vc.VideoServices.UpdateList(&video)
	if err != nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK, video) //return the user object
}

func (vc *VideoController) DeleteList(ctx *gin.Context) {
	title := ctx.Param("title")
	err := vc.VideoServices.DeleteList(&title)
	if err != nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// route function to group all the routes
func (vc *VideoController) RegisterVideoRoutes(rg *gin.RouterGroup) {
	videoRoute := rg.Group("/video") //base path

	videoRoute.GET("/get/:query/:page", vc.GetListByQuery)
	videoRoute.GET("/getAll", vc.GetAll)
	videoRoute.POST("/create", vc.CreateList)
	videoRoute.PATCH("/update", vc.UpdateList)
	videoRoute.DELETE("/delete/:query", vc.DeleteList)
}
