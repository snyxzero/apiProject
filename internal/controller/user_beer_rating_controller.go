package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/service/userbeerrating"
	"net/http"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type RatingRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id" binding:"required"`
	BeerID int `json:"beer_id" binding:"required"`
	Rating int `json:"rating" binding:"required"`
}

type UserBeerRatingController struct {
	repository   *repository.UserBeerRatingsRepository
	ratingPoints *userbeerrating.RatingPoints
}

func NewRatingController(repository *repository.UserBeerRatingsRepository, ratingPoints *userbeerrating.RatingPoints) *UserBeerRatingController {
	return &UserBeerRatingController{
		repository:   repository,
		ratingPoints: ratingPoints,
	}
}

func (o *UserBeerRatingController) GetRating(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	userBeerRating, err := o.repository.GetRating(c, id)
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

func (o *UserBeerRatingController) CreateRating(c *gin.Context) {
	var userBeerRatingRq RatingRequest
	err := c.ShouldBindJSON(&userBeerRatingRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	userBeerRating := models.UserBeerRating{
		User:   userBeerRatingRq.UserID,
		Beer:   userBeerRatingRq.BeerID,
		Rating: userBeerRatingRq.Rating,
	}

	userBeerRating, err = o.ratingPoints.AddRatingWithTransaction(c, &userBeerRating)
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

func (o *UserBeerRatingController) UpdateRating(c *gin.Context) {
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

	userBeerRating := models.UserBeerRating{
		ID:     id,
		User:   userBeerRatingRq.UserID,
		Beer:   userBeerRatingRq.BeerID,
		Rating: userBeerRatingRq.Rating,
	}

	userBeerRating, err = o.repository.UpdateRating(c, &userBeerRating)
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

func (o *UserBeerRatingController) DeleteRating(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = o.repository.DeleteRating(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
