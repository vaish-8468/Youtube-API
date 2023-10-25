//user implementation class

package services

import (
	"Youtube_RestAPI/api"
	"Youtube_RestAPI/models"
	"context"
	"errors"
	"log"
	"os"

	// "os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type VideoServiceImpl struct {
	//this will have user collection object which can be accessed using pointer
	videocollection *mongo.Collection 
	ctx             context.Context
}



// constructor
func NewVideoService(videocollection *mongo.Collection, ctx context.Context) VideoService {
	return &VideoServiceImpl{
		videocollection: videocollection,
		ctx:             ctx,
	}
}




//create receiver function to create function of type receiver
func (u *VideoServiceImpl) CreateList(video *models.Video) error {
	//logic for interacting with the database and create a new user

	//if there is an error occurred during user insertion, return it
	_, err := u.videocollection.InsertOne(u.ctx, video)
	return err
}



// receiver function to take video title as a parameter and return video object
func (u *VideoServiceImpl) GetList(title *string, page *string, pageSize *string) ([]*models.Video, error) {
	// //here, we have a dynamic variable i.e. video

	// Convert the page and pageSize to integers
	pageInt, _ := strconv.Atoi(*page)
	pageSizeInt, _ := strconv.Atoi(*pageSize)

	//Calculate the skip value for pagination
	skip := int64((pageInt - 1) * pageSizeInt)

	//define the max number of results to fetch
	maxResults := int64(20)

	//define filter to execute partial text search
	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: title}}}}


	// Define the sort order (descending by published datetime)
	sort := options.Find().SetSort(bson.D{{Key: "snippet.publishedAt", Value: -1}}).SetLimit(int64(pageSizeInt)).SetSkip(skip)

	// Find videos matching the search query
	cursor, err := u.videocollection.Find(u.ctx, filter, sort)

	//to store the videos
	var videos []*models.Video 

	if err != nil {

		panic(err)

	} else{

	foundResults:=false
	//check if data exists in the database, return it as a response

	for cursor.Next(context.TODO()) {
		foundResults= true //data exists in the database
		var video models.Video
		err := cursor.Decode(&video) 
		//we'll decode the query and assign it to the video variable
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}

	//stop the cursor
	cursor.Close(u.ctx)

	if !foundResults{
		// Data doesn't exist, call YouTube API to fetch data
		response, err := api.FetchYouTubeVideos(u.ctx, title, &maxResults, os.Getenv("APIkey"))
		if err != nil {
			log.Println(err)
		}

		// Iterate through each item in the response and store the relevant data.
		for _, video := range response.Items {

			// Create a Video instance for each video in the response
			publishedAt, err := time.Parse(time.RFC3339, video.Snippet.PublishedAt)
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
					ChannelTitle: video.Snippet.ChannelTitle,
				},
			}
			// Insert the video document into MongoDB
			_, err = u.videocollection.InsertOne(u.ctx, videos)
			if err != nil {
				log.Println(err)
			}
			//creating text indexing
			model := mongo.IndexModel{Keys: bson.D{{Key: "snippet.title", Value: "text"}, {Key: "snippet.description", Value: "text"}}}
			_, err = u.videocollection.Indexes().CreateOne(context.TODO(), model)
			if err != nil {
				panic(err)
			}
		}

		// Find videos matching the search query again after calling youtube api and storing the data
		cursor, err = u.videocollection.Find(u.ctx, filter, sort)
		if err != nil {
			panic(err)
		}
		//iterate through the cursor for storing the data
		for cursor.Next(context.TODO()) {
			var video models.Video
			err := cursor.Decode(&video) //we'll decode the query and assign it to the video variable
			if err != nil {
				return nil, err
			}
			videos = append(videos, &video)
	}

	//stop the cursor
	cursor.Close(u.ctx)
	}
}

	//return the list of videos fetched
	return videos, err
}




func (u *VideoServiceImpl) GetAll() ([]*models.Video, error) {
	//fetch videos one by one from the database
	var videos []*models.Video
	//so we'll use cursors
	cursors, err := u.videocollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	//otherwise, iterate through the cursors an use the next method to fetch each data
	for cursors.Next(u.ctx) {
		var video models.Video
		//decoding it an saving it in the user variable
		err := cursors.Decode(&video)

		if err != nil {
			return nil, err
		}
		//otherwise, append this video in the videos slice
		videos = append(videos, &video)

	}
	//traces error if occurred during the iteration
	if err := cursors.Err(); err != nil {
		return nil, err
	}

	//stop the cursor
	cursors.Close(u.ctx)

	if len(videos) == 0 {
		return nil, errors.New("no record found")
	}

	return videos, nil
}




func (u *VideoServiceImpl) UpdateList(video *models.Video) error {
	//filter query to find a user by filter name
	filter := bson.D{bson.E{Key: "title", Value: video.Snippet.Title}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "title", Value: video.Snippet.Title}, bson.E{Key: "description", Value: video.Snippet.Description}, bson.E{Key: "publishedAt", Value: video.Snippet.PublishedAt}, bson.E{Key: "etag", Value: video.Etag}}}} //we'll pass ab much object features as much they need to be changed
	result, _ := u.videocollection.UpdateOne(u.ctx, filter, update)

	//if there is a match count that means the user exist
	if result.MatchedCount != 1 {
		return errors.New("no matched record found for update")
	}
	return nil
}




func (u *VideoServiceImpl) DeleteList(title *string) error {
	//find the corresponding record to the title
	filter := bson.D{bson.E{Key: "title", Value: title}}
	result, _ := u.videocollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched record found to delete")
	}
	return nil
}
