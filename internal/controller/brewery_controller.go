package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"net/http"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type BreweryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type BreweryController struct {
	repository *repository.BreweriesRepository
}

func NewBreweryController(repository *repository.BreweriesRepository) *BreweryController {
	return &BreweryController{
		repository: repository,
	}
}

func (uc *BreweryController) GetBrewery(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	brewery, err := uc.repository.GetBrewery(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"brewery": brewery,
	})
	return
}

func (uc *BreweryController) CreateBrewery(c *gin.Context) {
	var breweryRq BreweryRequest
	err := c.ShouldBindJSON(&breweryRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	brewery := models.Brewery{
		Name: breweryRq.Name,
	}

	brewery, err = uc.repository.AddBrewery(c, &brewery)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"brewery": brewery,
	})

	return
}

func (uc *BreweryController) UpdateBrewery(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	var breweryRq BreweryRequest
	err = c.ShouldBindJSON(&breweryRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	brewery := models.Brewery{
		ID:   id,
		Name: breweryRq.Name,
	}

	brewery, err = uc.repository.UpdateBrewery(c, &brewery)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"brewery": brewery,
	})
	return
}

func (uc *BreweryController) DeleteBrewery(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = uc.repository.DeleteBrewery(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
