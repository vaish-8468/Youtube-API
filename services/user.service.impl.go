//user implementation class

package services

import (
	"FamPay/models"
	"context"
	"errors"
	"log"
	"strconv"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)





type UserServiceImpl struct{
	usercollection *mongo.Collection  //this will have user collection object which can be accessed using pointer
	ctx            context.Context
}

//constructor
func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService{
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx: ctx, //using context we can fix the execution time of any call
	}
}

//create receiver function to create function of type receiver

func (u *UserServiceImpl) CreateUser(user *models.User) error{
	//logic for interacting with the database and create a new user

	//if there is an error occurred during user insertion, return it
	_,err :=u.usercollection.InsertOne(u.ctx,user)
	return err
}

//receiver function to take username as a parameter and return user object
func (u *UserServiceImpl) GetUser(name *string, page *string, pageSize *string) ([]*models.User,error){
	// //here, we have a dynamic variable i.e. user
	// var user *models.User
	// //mongo query
	// query :=bson.D{bson.E{Key:"name", Value: name}}

	// //logic for finding user
	// err:= u.usercollection.FindOne(u.ctx,query).Decode(&user) //we'll decode the query and assign it to the user variable
	// return user, err

	// var user *models.User
	//mongo query
	// query := bson.D{bson.E{Key:"title", Value: title}}


	// check:=u.CreateIndex(name,true)
	// if check==false{
	// 	errors.New("Failed to create an index")
	// }
	
	 
 
	 // Convert the page and pageSize to integers
	 pageInt, _ := strconv.Atoi(*page)
	 pageSizeInt, _ := strconv.Atoi(*pageSize)
	 pageSizeInt64:=int64(pageSizeInt)
 
	 // Calculate the skip value for pagination
	 skip := int64((pageInt - 1) * pageSizeInt)

	 // Define the sort order (descending by published datetime)
	 sort := bson.D{{Key: "video.Snippet.PublishedAt", Value: -1}}

	 // Perform the MongoDB query with pagination and sorting
	 cursor, err := u.usercollection.Find(u.ctx, bson.M{}, &options.FindOptions{
		 Skip:  &skip,
		 Limit: &pageSizeInt64,
		 Sort:  sort,
	 })

	// filter := bson.M{
    //     "$text": bson.M{"$search": name},
    // }

    // Find videos matching the search query
    // cursor, err := u.usercollection.Find(u.ctx, filter)
    if err != nil {
        log.Println(err)
        // Handle the error and return an appropriate response
    }
    // defer cursor.Close(context.TODO())

    var users []*models.User
    for cursor.Next(context.TODO()) {
        var user models.User
        err:=cursor.Decode(&user)
		if err!=nil{
			return nil, err
		}
        users = append(users, &user)
    }
	//stop the cursor
	cursor.Close(u.ctx)
	
	if len(users)==0{
		 return nil, errors.New("no record found")
	}

	//logic for finding video
	// err:= u.videocollection.FindOne(u.ctx,query).Decode(&video) //we'll decode the query and assign it to the video variable
	return users, err



}


func (u *UserServiceImpl) GetAll() ([]*models.User,error){
	//fetch users one by one from the database
	var users []*models.User
	//so we'll use cursors
	cursors,err :=u.usercollection.Find(u.ctx,bson.D{{}})
	if err!=nil{
		return nil, err
	}
	//otherwise, iterate through the cursors an use the next method to fetch each data
	for cursors.Next(u.ctx){
		var user models.User 
		//decoding it an saving it in the user variable
		err:= cursors.Decode(&user)

		if err!=nil{
			return nil, err
		}
		//otherwise, append this user in the user slice
		users =append(users,&user)

	}
	//traces error if occurred during the iteration
	if err := cursors.Err(); err!=nil{
		return nil, err
	}

	//stop the cursor
	cursors.Close(u.ctx)

	if len(users)==0{
		 return nil, errors.New("no record found")
	}

	return users,nil
}


func (u *UserServiceImpl) UpdateUser(user *models.User) error{
	//filter query to find a user by filter name
	filter := bson.D{bson.E{Key:"user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_address", Value:user.Age}, bson.E{Key: "user_age", Value: user.Address}}}} //we'll pass ab much object features as much they need to be changed
	result,_ := u.usercollection.UpdateOne(u.ctx,filter,update)

	//if there is a match count that means the user exist
	if result.MatchedCount!=1{
		return errors.New("no matched document found for update")
	}
	return nil
}



func (u *UserServiceImpl) DeleteUser(name *string) error{
	//find the corresponding record to the name
	filter := bson.D{bson.E{Key:"user_name", Value: name}}
	result,_ := u.usercollection.DeleteOne(u.ctx,filter)
	if result.DeletedCount!=1{
		return errors.New("no matched document found to delete")
	}
	return nil
}

// func (u *UserServiceImpl) CreateIndex(field *string, unique bool) bool {

//     // 1. Lets define the keys for the index we want to create
   
// 	mod  := mongo.IndexModel{
// 		Keys: bson.D{{Key: *field, Value: "text"}},
// 		Options: options.Index().SetUnique(unique),
// 	}

//     // 2. Create the context for this operation
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

//     // 3. Connect to the database and access the collection
//     collection := u.usercollection;

//     // 4. Create a single index
// 	_, err := collection.Indexes().CreateOne(ctx, mod)
// 	if err != nil {
//         // 5. Something went wrong, we log it and return false
// 		log.Println(err)
// 		return false
//     }

//     // 6. All went well, we return true
// 	return true
// }