package models
import "go.mongodb.org/mongo-driver/bson/primitive"


// // struct to store video metadata
// type Video struct {
// 	Id           primitive.ObjectID `json:"_id" bson:"_id"`
// 	Title        string             `json:"title" bson:"title"`
// 	Description  string             `json:"description" bson:"description"`
// 	PublishedAt  string             `json:"published_at" bson:"published_at"`
// 	ThumbnailUrl string             `json:"thumbnail_url" bson:"thumbnail_url"`
// 	VideoETag    string             `json:"video_etag" bson:"video_etag"`
// }

// struct to store data from user
type ShowVideoRequest struct {
	Page        int64  `json:"page"`
	SearchQuery string `json:"search_query"`
}



type Video struct {
	Kind string `json:"kind" bson:"kind"`
	Etag string `json:"etag" bson:"etag"`
	Id Id  `json:"id" bson:"id"`
	Snippet Snippet `json:"snippet" bson:"snippet"`
	
}

type Id struct{
	Kind string `json:"kind" bson:"kind"`
	VideoId string `json:"videoId" bson:"videoId"`
	ChannelId string `json:"channelId" bson:"channelId"`
	PlaylistId string `json:"playlistId" bson:"playlistId"`
} 

type Snippet struct{
	PublishedAt primitive.DateTime `json:"publishedAt" bson:"publishedAt"`
	ChannelId string `json:"channelId" bson:"channelId"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	ThumbnailUrl          map[string]struct {
		URL    string `json:"url"`
		Width  uint   `json:"width"`
		Height uint   `json:"height"`
	} `json:"thumbnails" bson:"thumbnails"`
	ChannelTitle  string `json:"channelTitle" bson:"channelTitle"`
	LiveBroadCastContent string `json:"liveBroadcastContent" bson:"liveBroadcastContent"`

}

