//user implementation class

package services

import (
	"FamPay/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoServiceImpl struct{
	videocollection *mongo.Collection  //this will have user collection object which can be accessed using pointer
	ctx            context.Context
}


//constructor
func NewVideoService(videocollection *mongo.Collection, ctx context.Context) VideoService{
	return &VideoServiceImpl{
		videocollection: videocollection,
		ctx: ctx,
	}
}

//create receiver function to create function of type receiver

func (u *VideoServiceImpl) CreateList(video *models.Video) error{
	//logic for interacting with the database and create a new user

	//if there is an error occurred during user insertion, return it
	_,err :=u.videocollection.InsertOne(u.ctx,video)
	return err
}

//receiver function to take video title as a parameter and return video object
func (u *VideoServiceImpl) GetList(title *string) (*models.Video,error){
	//here, we have a dynamic variable i.e. video
	var video *models.Video
	//mongo query
	query := bson.D{bson.E{Key:"title", Value: title}}

	//logic for finding video
	err:= u.videocollection.FindOne(u.ctx,query).Decode(&video) //we'll decode the query and assign it to the video variable
	return video, err
}


func (u *VideoServiceImpl) GetAll() ([]*models.Video,error){
		//fetch videos one by one from the database
		var videos []*models.Video
		//so we'll use cursors
		cursors,err :=u.videocollection.Find(u.ctx,bson.D{{}})
		if err!=nil{
			return nil, err
		}
		//otherwise, iterate through the cursors an use the next method to fetch each data
		for cursors.Next(u.ctx){
			var video models.Video 
			//decoding it an saving it in the user variable
			err:= cursors.Decode(&video)
	
			if err!=nil{
				return nil, err
			}
			//otherwise, append this video in the videos slice
			videos =append(videos,&video)
	
		}
		//traces error if occurred during the iteration
		if err := cursors.Err(); err!=nil{
			return nil, err
		}
	
		//stop the cursor
		cursors.Close(u.ctx)
	
		if len(videos)==0{
			 return nil, errors.New("no record found")
		}
	
		return videos,nil
	}



func (u *VideoServiceImpl) UpdateList(video *models.Video) error{
	//filter query to find a user by filter name
	filter := bson.D{bson.E{Key:"title", Value: video.Snippet.Title}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "title", Value: video.Snippet.Title}, bson.E{Key: "description", Value:video.Snippet.Description}, bson.E{Key: "thumbnails", Value: video.Snippet.ThumbnailUrl}, bson.E{Key: "publishedAt", Value:video.Snippet.PublishedAt}, bson.E{Key: "etag", Value:video.Etag}}}} //we'll pass ab much object features as much they need to be changed
	result,_ := u.videocollection.UpdateOne(u.ctx,filter,update)

	//if there is a match count that means the user exist
	if result.MatchedCount!=1{
		return errors.New("no matched record found for update")
	}
	return nil
}



func (u *VideoServiceImpl) DeleteList(title *string) error{
	//find the corresponding record to the title
	filter := bson.D{bson.E{Key:"title", Value: title}}
	result,_ := u.videocollection.DeleteOne(u.ctx,filter)
	if result.DeletedCount!=1{
		return errors.New("no matched record found to delete")
	}
	return nil
}