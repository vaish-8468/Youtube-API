### Project Goal
To make an API to fetch latest videos sorted in reverse chronological order of their
publishing date-time from YouTube for a given tag/search query in a paginated
response.

Basic Requirements:

• Server should call the YouTube API continuously in background (async) with
some interval (say 10 seconds) for fetching the latest videos for a predefined
search query and should store the data of videos (specifically these fields -
Video title, description, publishing datetime, thumbnails URLs and any other
fields you require) in a database with proper indexes.

• A `GET` API which returns the stored video data in a paginated response sorted
in descending order of published datetime.

• A basic search API to search the stored videos using their title and description.

• Dockerize the project.

• It should be scalable and optimised.

### Tech Stack Used
1. Golang's Gin-Gonic Web Framework
2. MongoDB driver

### Methodology
This projected has been implemented using Golang's popular Gin Web Framework and used MongoDB driver to perform CRUD operations created using MVC([Model-View-Controller](https://www.geeksforgeeks.org/mvc-framework-introduction/))framework.

![Model1](https://github.com/vaish-8468/Youtube-API/assets/84587662/f7d997ea-6097-4b7f-8660-f4ffe416de63)


### APIs
`/Get` Returns list of videos paginated with 10 items per page.
```
localhost:9090/Get?page=1
```

`/Get/title` Returns list of videos with partially or completely match the given title query parameter, paginated with 10 items per page.
```
localhost:9090/search?page=1&query=DBMS
```

### Database
This server uses MongoDB. To handle search(query) queries, we have 2 text indexes (compound index) on title and description fields.

### Usage
Clone the repository using :
```
git clone https://github.com/vaish-8468/Youtube-API.git
```

### Reference:
 1. [YouTube data v3 API](https://developers.google.com/youtube/v3/gettingstarted)
 2. [Search API reference](https://developers.google.com/youtube/v3/docs/search/list)
 3. [Go REST API – Sort, Page, Filter](https://go-cloud-native.com/golang/go-rest-api-sort-page-and-filter)
 4. [Create a Text Index](https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/text/#:~:text=To%20perform%20a%20text%20search,field%20in%20your%20query%20filter)



