package controller

import (
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

func ValidID(idParam string) (int, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errorcrud.ErrInvalidFormat
	}
	if id < 1 {
		return 0, errorcrud.ErrNegativeID
	}
	return id, nil
}

type BeerRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	BreweriesID int    `json:"breweries_id" binding:"required"`
}

type BeerController struct {
	repository *repository.BeersRepository
}

func NewBeerController(repository *repository.BeersRepository) *BeerController {
	return &BeerController{
		repository: repository,
	}
}

func (o *BeerController) GetBeer(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	beer, err := o.repository.GetBeer(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   beer,
	})
	return
}

func (o *BeerController) CreateBeer(c *gin.Context) {

	var beerRq BeerRequest
	err := c.ShouldBindJSON(&beerRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	beer := models.Beer{
		Name:    beerRq.Name,
		Brewery: beerRq.BreweriesID,
	}

	beer, err = o.repository.AddBeer(c, &beer)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   beer,
	})
	return
}

func (o *BeerController) UpdateBeer(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	var beerRq BeerRequest
	err = c.ShouldBindJSON(&beerRq)
	if err != nil {
		errorcrud.ErrInvalidJson(c, err)
		return
	}

	beer := models.Beer{
		ID:      id,
		Name:    beerRq.Name,
		Brewery: beerRq.BreweriesID,
	}

	beer, err = o.repository.UpdateBeer(c, &beer)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   beer,
	})
	return
}

func (o *BeerController) DeleteBeer(c *gin.Context) {
	id, err := ValidID(c.Param("id"))
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	err = o.repository.DeleteBeer(c, id)
	if err != nil {
		errorcrud.ErrorCheck(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//
