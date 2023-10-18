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

### Functionalities
1. The server spawns a go routine which gets videos metadata (with predefined query, "DBMS" in our case) from youtube every 10seconds.
2. User can supply multiple API keys, first valid API key in the list will be used everytime a request is made.
3. Search query matches with objects with partially or completely matching title or description. The search is case insensitive.
4. User can retrieve the latest videos in reverse chronological order of their published time
5. User can perform CRUD operations in mongodb database as per the requirements.


### Methodology
This projected has been implemented using Golang's popular Gin Web Framework and used MongoDB driver to perform CRUD operations created using MVC([Model-View-Controller](https://www.geeksforgeeks.org/mvc-framework-introduction/))framework.

![Model1](https://github.com/vaish-8468/Youtube-API/assets/84587662/3d8a22a3-2fe9-43f1-8ce9-27f6ee95bcef)



### APIs
Routes:
![image](https://github.com/vaish-8468/Youtube-API/assets/84587662/7c06824d-6e02-4216-bb11-3b1267cfbe54)


`/Get` Returns list of videos paginated with 10 items per page.
```
localhost:9090/v2/video/get?page=1
```

`/Get/:query` Returns list of videos with partially or completely match the given title query parameter, paginated with 10 items per page.
```
localhost:9090/v2/video/get?page=1&query=DBMS
```
![image](https://github.com/vaish-8468/Youtube-API/assets/84587662/bb5bbb79-64dc-4023-b685-0c2c45095909)



### Database
This server uses MongoDB. To handle search(query) queries, we have 2 text indexes (compound index) on title and description fields.

### Usage
Clone the repository using :
```
git clone https://github.com/vaish-8468/Youtube-API.git
```
To start the server:
1. `make run` will start the server locally on port 9090
   or
`docker build . -t docker-gs-ping && docker run -p 9090:9090 docker-gs-ping` will start a docker container on port 9090.
![image](https://github.com/vaish-8468/Youtube-API/assets/84587662/58094105-ef0e-4e75-bc34-c3cb059e222b)


### Issues
To fix mongodb error, we can start mongodb server manually:
```
sudo systemctl start mongod
mongosh
```
![image](https://github.com/vaish-8468/Youtube-API/assets/84587662/3814e17e-4dbd-4ffc-b0c5-e15af7a73488)



### Reference:
 1. [YouTube data v3 API](https://developers.google.com/youtube/v3/gettingstarted)
 2. [Search API reference](https://developers.google.com/youtube/v3/docs/search/list)
 3. [Go REST API – Sort, Page, Filter](https://go-cloud-native.com/golang/go-rest-api-sort-page-and-filter)
 4. [Create a Text Index](https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/text/#:~:text=To%20perform%20a%20text%20search,field%20in%20your%20query%20filter)
 5. [Go Docker Image](https://docs.docker.com/language/golang/build-images/)



