
// Generally, API contracts describe the supported responses, supported methods, such as POST and PUT, etc., terms of service, and version, along with the outputs and inputs. 


//services interact with database
package services

import "FamPay/models"


type UserService interface{
	CreateUser(*models.User) error  //will save the error if reflected in mongodb
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)  //returns all the objects as slice
	UpdateUser(*models.User) error  //takes user information as a parameter
	DeleteUser(*string) error //takes username as a parameter

	
}