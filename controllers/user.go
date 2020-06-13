package controllers

import (
	"github.com/gin-gonic/gin"
	// "github.com/hepiska/todo-go/models/entity"

	"github.com/hepiska/todo-go/models/entity"
	"github.com/hepiska/todo-go/models/service"
)

type queryStruct struct {
	Search string `form:"search"`
	Skip   int64  `form:"skip"`
	Limit  int64  `form:"limit"`
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

	users, count, err := userService.Find(query.Search, query.Skip, query.Limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "system error"})
		return
	}

	c.JSON(200, gin.H{"data": users, "total": count})
	return

}

// GetOne user
func (userC *UserController) GetOne(c *gin.Context) {
	_id := c.Param("id")

	userService := service.Userservice{}

	user, err := userService.FindOneByID(_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": user})
	return

}

// Delete delete user
func (userC *UserController) Delete(c *gin.Context) {
	_id := c.Param("id")
	userService := service.Userservice{}
	err := userService.Delete(_id)
	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(500, gin.H{"status": "ok"})
	return

}

// Update  user
func (userC *UserController) Update(c *gin.Context) {
	_id := c.Param("id")

	input := entity.UserEdit{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(401, gin.H{"error": "Please input valid data asas"})
		return
	}
	userService := service.Userservice{}
	errEdit := userService.Update(_id, &input)
	if errEdit != nil {
		c.JSON(500, gin.H{"error": errEdit.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
	return

}
