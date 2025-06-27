package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/service"
	"net/http"

	"github.com/snyxzero/apiProject/internal/models"
)

type BreweryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type BreweryController struct {
	service *service.BreweryService
}

func NewBreweryController(service *service.BreweryService) *BreweryController {
	return &BreweryController{
		service: service,
	}
}

func (o *BreweryController) GetBrewery(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}
	var brewery *models.Brewery
	brewery, err = o.service.GetBrewery(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   brewery,
	})
	return
}

func (o *BreweryController) CreateBrewery(c *gin.Context) {
	var breweryRq BreweryRequest
	err := c.ShouldBindJSON(&breweryRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	brewery := &models.Brewery{
		Name: breweryRq.Name,
	}

	brewery, err = o.service.AddBrewery(c, brewery)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   brewery,
	})

	return
}

func (o *BreweryController) UpdateBrewery(c *gin.Context) {
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

	brewery := &models.Brewery{
		ID:   id,
		Name: breweryRq.Name,
	}

	brewery, err = o.service.UpdateBrewery(c, brewery)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   brewery,
	})
	return
}

func (o *BreweryController) DeleteBrewery(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = o.service.DeleteBrewery(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
