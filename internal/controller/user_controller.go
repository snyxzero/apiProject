package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/service"
	"net/http"
)

type UserRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	RatingPoints int
}

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (o *UserController) GetUser(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	user, err := o.service.GetUser(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
	return
}

func (o *UserController) CreateUser(c *gin.Context) {
	var userRq UserRequest
	err := c.ShouldBindJSON(&userRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	user := &models.User{
		Name: userRq.Name,
	}

	user, err = o.service.AddUser(c, user)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
	return
}

func (o *UserController) UpdateUser(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	var userRq UserRequest
	err = c.ShouldBindJSON(&userRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	user := &models.User{
		ID:   id,
		Name: userRq.Name,
	}

	user, err = o.service.UpdateUser(c, user)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
	return
}

func (o *UserController) DeleteUser(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = o.service.DeleteUser(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
