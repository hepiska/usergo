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
