package service

import (
	"errors"
	"fmt"

	"github.com/hepiska/todo-go/models/db"
	"github.com/hepiska/todo-go/models/entity"

	"github.com/goonode/mogo"
	"labix.org/v2/mgo/bson"
)

type Userservice struct{}

//Create is to create  new user
func (userservice Userservice) Create(user *(entity.User)) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	fmt.Println("pre item is ", doc)
	if err == nil {
		return errors.New("Already Exist")
	}
	userModel := mogo.NewDoc(user).(*(entity.User))
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

// find one user
func (userservice Userservice) FindOne(user *entity.User) (*(entity.User), error) {
	conn := db.GetConnection()
	defer conn.Session.Close()
	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)

	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Find to find many
func (userservice Userservice) Find(search string, skip int, limit int) ([]entity.User, error) {
	conn := db.GetConnection()

	defer conn.Session.Close()

	doc := mogo.NewDoc(entity.User{}).(*(entity.User))
	users := []entity.User{}
	user := &entity.User{}
	condition := bson.M{"$or": []bson.M{bson.M{"email": bson.M{"$regex": search}}, bson.M{"name": bson.M{"$regex": search}}}}
	iter := doc.Find(condition).Limit(limit).Skip(skip * limit).Iter()

	for iter.Next(user) {
		users = append(users, *user)
	}

	return users, nil

}

// find by email
func (userservice Userservice) FindbyEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	user.Email = email
	return userservice.FindOne(user)

}
