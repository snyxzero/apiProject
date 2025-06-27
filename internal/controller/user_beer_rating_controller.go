package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/service"
	"net/http"
)

type RatingRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id" binding:"required"`
	BeerID int `json:"beer_id" binding:"required"`
	Rating int `json:"rating" binding:"required"`
}

type UserBeerRatingController struct {
	service      *service.UserBeerRatingService
	ratingPoints *service.RatingPoints
}

func NewUserBeerRatingController(service *service.UserBeerRatingService) *UserBeerRatingController {
	return &UserBeerRatingController{
		service: service,
	}
}

func (o *UserBeerRatingController) GetUserBeerRating(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	userBeerRating, err := o.service.GetUserBeerRating(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"rating": userBeerRating,
	})
	return
}

func (o *UserBeerRatingController) CreateUserBeerRating(c *gin.Context) {
	var userBeerRatingRq RatingRequest
	err := c.ShouldBindJSON(&userBeerRatingRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	userBeerRating := &models.UserBeerRating{
		User:   userBeerRatingRq.UserID,
		Beer:   userBeerRatingRq.BeerID,
		Rating: userBeerRatingRq.Rating,
	}

	userBeerRating, err = o.service.AddUserBeerRatingWithTransaction(c, userBeerRating)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         "success",
		"userBeerRating": userBeerRating,
	})
	return
}

func (o *UserBeerRatingController) UpdateUserBeerRating(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	var userBeerRatingRq RatingRequest
	err = c.ShouldBindJSON(&userBeerRatingRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	userBeerRating := &models.UserBeerRating{
		ID:     id,
		User:   userBeerRatingRq.UserID,
		Beer:   userBeerRatingRq.BeerID,
		Rating: userBeerRatingRq.Rating,
	}

	userBeerRating, err = o.service.UpdateUserBeerRating(c, userBeerRating)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   userBeerRating,
	})
	return
}

func (o *UserBeerRatingController) DeleteUserBeerRating(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = o.service.DeleteUserBeerRating(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
