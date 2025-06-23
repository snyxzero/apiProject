package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type ratingClipboard struct {
	ID     int `json:"id"`
	User   int `json:"user" binding:"required"` // должно быть user_id и для пива так же
	Beer   int `json:"beer" binding:"required"`
	Rating int `json:"rating" binding:"required"`
}

// название должно совпадать полностью, для читаемости и понимамости, то есть UserBeerRatingCOntroller
type RatingController struct {
	repository *repository.UserBeerRatingRepository
}

func NewRatingController(repository *repository.UserBeerRatingRepository) *RatingController {
	return &RatingController{
		repository: repository,
	}
}

func (o *RatingController) GetRating(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if id < 1 {
		log.Println("incorrect id (id < 1)")
		c.Status(http.StatusBadRequest)
		return
	}
	userBeerRating, err := o.repository.GetRating(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"rating": userBeerRating,
	})
	return
}

func (o *RatingController) CreateRating(c *gin.Context) {

	var ratingCb ratingClipboard
	err := c.ShouldBindJSON(&ratingCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	rating := models.UserBeerRating{
		User:   ratingCb.User,
		Beer:   ratingCb.Beer,
		Rating: ratingCb.Rating,
	}

	ratingCb.ID, err = o.repository.AddRating(c, rating)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"rating": ratingCb,
	})
	return
}

func (o *RatingController) UpdateRating(c *gin.Context) {

	var ratingCb ratingClipboard
	err := c.ShouldBindJSON(&ratingCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	rating := models.UserBeerRating{
		User:   ratingCb.User,
		Beer:   ratingCb.Beer,
		Rating: ratingCb.Rating,
	}

	err = o.repository.UpdateRating(c, rating)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   ratingCb,
	})
	return
}

func (o *RatingController) DeleteRating(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if id < 1 {
		log.Println("incorrect id (id < 1)")
		c.Status(http.StatusBadRequest)
		return
	}
	err = o.repository.DeleteRating(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
