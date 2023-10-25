
// Generally, API contracts describe the supported responses, supported methods, such as POST and PUT, etc., terms of service, and version, along with the outputs and inputs. 

package services

import "Youtube_RestAPI/models"


type VideoService interface{
	CreateList(*models.Video) error  //will save the error if reflected in mongodb
	GetList(*string,*string,*string) ([]*models.Video, error)
	GetAll() ([]*models.Video, error)  //returns all the objects as slice
	UpdateList(*models.Video) error  //takes video information as a parameter
	DeleteList(*string) error //takes key as a parameter


}