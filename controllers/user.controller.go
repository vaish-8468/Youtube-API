package controllers

import (
	"FamPay/models"
	"FamPay/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//interacts with user service

type UserController struct{
	UserServices services.UserService //reference of userservice
}

func New(userservice services.UserService) UserController{
	return UserController{
		UserServices: userservice,

	}
}

//from controller we'll define routes , hence we'll call those methods defines in models
//gin.context holds information about the request that we're gonna send and will get a json object response
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&user); err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	//we'll check for the error while creating new user
	err := uc.UserServices.CreateUser(&user) //passing the address of the user object
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

//route functions or handlers
func (uc *UserController) GetUser(ctx *gin.Context){
	var username string =ctx.Param("name") //fetching the name from ctx object
	user, err := uc.UserServices.GetUser(&username)
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK,user) //return the user object
}


func (uc *UserController) GetAll(ctx *gin.Context){
	users, err := uc.UserServices.GetAll()
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	ctx.JSON(http.StatusOK,users) //return the user object
}


func (uc *UserController) UpdateUser(ctx *gin.Context){
	var user models.User
	//we'll check if there is any error while binding to the user variable
	if err := ctx.ShouldBindJSON(&user); err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}
	err := uc.UserServices.UpdateUser(&user)
	if err!=nil {
		//error while saving in the mongodb
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
	}

	ctx.JSON(http.StatusOK,user) //return the user object
}



func (uc *UserController) DeleteUser(ctx *gin.Context){
	username := ctx.Param("name")
	err := uc.UserServices.DeleteUser(&username)
	if err!=nil{
			//error while saving in the mongodb
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) //returns the message of httpBadRequest output
		}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}


//route function to grou all the routes
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userRoute :=rg.Group("/user")  //base path
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getAll", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}