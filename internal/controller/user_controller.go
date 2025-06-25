package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
	"net/http"
)

type UserRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	RatingPoints int
}

type UserController struct {
	repository *repository.UsersRepository
}

func NewUserController(repository *repository.UsersRepository) *UserController {
	return &UserController{
		repository: repository,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	user, err := uc.repository.GetUser(c, id)
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

func (uc *UserController) CreateUser(c *gin.Context) {
	var userRq UserRequest
	err := c.ShouldBindJSON(&userRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	user := models.User{
		Name: userRq.Name,
	}

	user, err = uc.repository.AddUser(c, &user)
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

func (uc *UserController) UpdateUser(c *gin.Context) {
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

	user := models.User{
		ID:   id,
		Name: userRq.Name,
	}

	user, err = uc.repository.UpdateUser(c, &user)
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

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = uc.repository.DeleteUser(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
