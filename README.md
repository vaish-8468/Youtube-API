# YOUTUBE-GO RestAPI
### Content
1. [Project Goal](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#project-goal)
2. [TechStack Used](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#tech-stack-used)
3. [Functionalities](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#functionalities)
4. [Methodology](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#methodology)
5. [APIs](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#apis)
6. [Database](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#database)
7. [Usage](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#usage)
8. [References](https://github.com/vaish-8468/Youtube-API/edit/main/README.md#reference)


Project Directory:
```
Youtbe_RestAPI
|---models
|     |---video.go
|
|---configs
|     |---mongodb.go
|
|----api
|     |---youtube.go
|
|
|---services
|     |---video.services.go
|     |---video.services.impl.go
|
|---controllers
|     |---video.controller.go
|
|---go.mod
|---go.sum
|---README.md
|---Dockerfile
|---main.go

```


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
2. Search query matches with objects with partially or completely matching title or description. The search is case insensitive.
3. User can retrieve the latest videos in reverse chronological order of their published time
4. User can perform CRUD operations in mongodb database as per the requirements.


### Methodology
This projected has been implemented using Golang's popular Gin Web Framework and used MongoDB driver to perform CRUD operations created using MVC([Model-View-Controller](https://www.geeksforgeeks.org/mvc-framework-introduction/))framework.

![image](https://github.com/vaish-8468/Youtube_RestAPI/assets/84587662/86b53b56-9167-485c-8e6b-2524e5435d07)



User Flow Diagram:

![image](https://github.com/vaish-8468/Youtube_RestAPI/assets/84587662/3d25c59b-c12d-436e-8a07-c91f323671ee)




### APIs
Routes:
![image](https://github.com/vaish-8468/Youtube_RestAPI/assets/84587662/f20d9d85-3447-4963-ab81-b16550555693)



`/Get/:query/:page` Returns list of videos with partially or completely match the given title query parameter, paginated with 10 items per page.
```
localhost:9090/v2/video/get/dbms/2
```
<img width="1440" alt="image" src="https://github.com/vaish-8468/Youtube_RestAPI/assets/84587662/b6e08811-4a9e-4289-a701-d8e6c268ffcb">





### Database
This server uses MongoDB. 

To handle search(query) queries, we have 2 text indexes (compound index) on title and description fields.For more details, [refer to](https://www.mongodb.com/docs/drivers/go/current/fundamentals/connection/).


### Usage
Clone the repository using :
```
git clone https://github.com/vaish-8468/Youtube_RestAPI.git
```
To start the server:
1. `make run` will start the server locally on port 9090
   or
`docker build . -t docker-gs-ping && docker run -p 9090:9090 docker-gs-ping` will start a docker container on port 9090.
<img width="945" alt="image" src="https://github.com/vaish-8468/Youtube_RestAPI/assets/84587662/706dfc77-ec58-4547-ae45-0dbdfd52fc61">






### Reference:
 1. [YouTube data v3 API](https://developers.google.com/youtube/v3/gettingstarted)
 2. [Search API reference](https://developers.google.com/youtube/v3/docs/search/list)
 3. [Go REST API – Sort, Page, Filter](https://go-cloud-native.com/golang/go-rest-api-sort-page-and-filter)
 4. [Create a Text Index](https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/text/#:~:text=To%20perform%20a%20text%20search,field%20in%20your%20query%20filter)
 5. [Go Docker Image](https://docs.docker.com/language/golang/build-images/)



