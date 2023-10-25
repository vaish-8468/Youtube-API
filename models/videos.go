package models

import "time"

// import "go.mongodb.org/mongo-driver/bson/primitive"

// // struct to store video metadata

type ShowVideoRequest struct {
	Page        int64  `json:"page"`
	SearchQuery string `json:"search_query"`
}

type Video struct {
	Kind    string  `json:"kind" bson:"kind"`
	Etag    string  `json:"etag" bson:"etag"`
	Id      Id      `json:"id" bson:"id"`
	Snippet Snippet `json:"snippet" bson:"snippet"`
}

type Id struct {
	IDKind     string `json:"kind" bson:"kind"`
	VideoId    string `json:"videoId" bson:"videoId"`
	ChannelId  string `json:"channelId" bson:"channelId"`
	PlaylistId string `json:"playlistId" bson:"playlistId"`
}

type Snippet struct {
	PublishedAt      time.Time `json:"publishedAt" bson:"publishedAt"`
	SnippetChannelId string `json:"channelId" bson:"channelId"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	ChannelTitle         string   `json:"channelTitle" bson:"channelTitle"`

}
