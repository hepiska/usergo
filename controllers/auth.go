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
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
		Address  string `json:"address"`
	}
	var info signupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(401, gin.H{"error": "Please input valid data asas"})
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

// Login sudah jelas
func (auth *AuthController) Login(c *gin.Context) {

	var info entity.User
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	userservice := service.Userservice{}
	user, err := userservice.FindbyEmail(info.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password)); err != nil {
		c.JSON(401, gin.H{"error": "email and password invalid"})
		return
	}

	token, errToken := user.GetJwtToken()
	if errToken != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"token": token})
}

// GetMe get user data base on token
func (auth *AuthController) GetMe(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	c.JSON(200, gin.H{"name": user.Name, "email": user.Email, "address": user.Address})
}
