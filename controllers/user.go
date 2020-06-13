package controllers

import (
	"github.com/gin-gonic/gin"
	// "github.com/hepiska/todo-go/models/entity"
	"github.com/hepiska/todo-go/models/service"
)

type queryStruct struct {
	Search string `form:"search"`
	Skip   int    `form:"skip"`
	Limit  int    `form:"limit"`
}

// UserController for user
type UserController struct{}

// GetAll user
func (userC *UserController) GetAll(c *gin.Context) {
	var query queryStruct
	query.Limit = 10
	query.Search = ""
	query.Skip = 0
	c.ShouldBindQuery(&query)

	userService := service.Userservice{}

	users, err := userService.Find(query.Search, query.Skip, query.Limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "system error"})
		return
	}

	c.JSON(200, gin.H{"data": users})
	return

}
