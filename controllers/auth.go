package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hepiska/todo-go/models/entity"
	"github.com/hepiska/todo-go/models/service"
	"golang.org/x/crypto/bcrypt"
)

//AuthController is for auth logic
type AuthController struct{}

//Signup is for user signup
func (auth *AuthController) Signup(c *gin.Context) {

	type signupInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Address  string `json:"address"`
	}
	var info signupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	user := entity.User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return

	}

	user.Password = string(hash)
	user.Name = info.Name
	user.Address = info.Address
	userService := service.Userservice{}
	err = userService.Create(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetJwtToken()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})

	return
}
