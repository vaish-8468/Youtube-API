//user implementation class

package services

import (
	"FamPay/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
func (u *UserServiceImpl) GetUser(name *string) (*models.User,error){
	//here, we have a dynamic variable i.e. user
	var user *models.User
	//mongo query
	query :=bson.D{bson.E{Key:"name", Value: name}}

	//logic for finding user
	err:= u.usercollection.FindOne(u.ctx,query).Decode(&user) //we;'' decode the query and assign it to the user variable
	return user, err
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