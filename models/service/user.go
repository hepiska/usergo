package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hepiska/todo-go/models/db"
	"github.com/hepiska/todo-go/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Userservice struct{}

// Create is to create  new user
func (userservice Userservice) Create(user *(entity.User)) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	db := db.ConfigDB()
	// var user entity.User

	count, err := db.Collection("users").CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email being used")
	}
	_, errIns := db.Collection("users").InsertOne(ctx, user)
	if errIns != nil {
		return errIns
	}

	return nil
}

// FindOneByID user
func (userservice Userservice) FindOneByID(id string) (*(entity.User), error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	oid, erroid := primitive.ObjectIDFromHex(id)
	if erroid != nil {
		return nil, erroid
	}

	db := db.ConfigDB()
	var user entity.User
	err := db.Collection("users").FindOne(ctx, bson.M{"_id": oid}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete delete user by id
func (userservice Userservice) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	oid, erroid := primitive.ObjectIDFromHex(id)
	if erroid != nil {
		return erroid
	}

	db := db.ConfigDB()
	err := db.Collection("users").FindOneAndDelete(ctx, bson.M{"_id": oid}).Err()
	if err != nil {
		return err
	}
	return nil

}

// // Update delete user by id
// func (userservice Userservice) Update(_id string, user *entity.User) error {
// 	conn := db.GetConnection()
// 	defer conn.Session.Close()

// 	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
// 	err := doc.FindOne(bson.M{"_id": _id}, doc)
// 	if err != nil {
// 		return err
// 	}

// 	doc.Address = user.Address

// 	doc.Name = user.Name

// 	errsave := doc.Save()

// 	if errsave != nil {
// 		return err
// 	}

// 	return nil

// }

// Find to find many
func (userservice Userservice) Find(search string, skip int64, limit int64) ([]entity.User, int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db := db.ConfigDB()

	users := []entity.User{}
	var user entity.User

	condition := bson.M{"$or": []bson.M{
		bson.M{"email": bson.M{"$regex": search}},
	}}
	findOptions := options.Find()
	findOptions.SetSkip(skip * limit)
	findOptions.SetLimit(limit)
	findOptions.SetProjection(bson.M{"password": 0})

	curr, err := db.Collection("users").Find(ctx, condition, findOptions)
	count, errcount := db.Collection("users").CountDocuments(ctx, condition)

	fmt.Println("count", count)

	if errcount != nil {
		return nil, 0, errcount
	}
	if err != nil {
		return nil, 0, err
	}

	for curr.Next(ctx) {
		errdecode := curr.Decode(&user)
		if errdecode != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		users = append(users, user)
	}

	return users, count, nil

}

// FindbyEmail find by email
func (userservice Userservice) FindbyEmail(email string) (*entity.User, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	db := db.ConfigDB()

	var user entity.User

	err := db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}
	fmt.Print(user)

	return &user, nil

}
