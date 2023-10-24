package api

import (
	"context"
	"log"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchYouTubeVideos(ctx context.Context, query *string, maxResults *int64, developerKey string) (*youtube.SearchListResponse, error) {

    // Initialize the YouTube service with the API key using the "option" package.
    service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
    if err != nil {
        log.Printf("Error creating a new YouTube client: %v\n", err)
        return nil, err
    }

    // Make the API call to YouTube.
    call := service.Search.List([]string{"id", "snippet"}).
        Q(*query).
        MaxResults(*maxResults)

    response, err := call.Do()
    if err != nil {
        log.Println("Error encountered:", err)
        return nil, err
    }

    return response, nil
}


