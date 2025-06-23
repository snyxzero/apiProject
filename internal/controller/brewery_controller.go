package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type breweryClipboard struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type BreweryController struct {
	repository *repository.BreweryRepository
}

func NewBreweryController(repository *repository.BreweryRepository) *BreweryController {
	return &BreweryController{
		repository: repository,
	}
}

func (uc *BreweryController) GetBrewery(c *gin.Context) {
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
	brewery, err := uc.repository.GetBrewery(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"brewery": brewery,
	})
	return
}

func (uc *BreweryController) CreateBrewery(c *gin.Context) {

	var breweryCb breweryClipboard
	err := c.ShouldBindJSON(&breweryCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	user := models.Brewery{ // название не user)
		Name: breweryCb.Name,
	}

	breweryCb.ID, err = uc.repository.AddBrewery(c, user)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  " brewery create",
		"brewery": breweryCb,
	})

	return
}

func (uc *BreweryController) UpdateBrewery(c *gin.Context) {

	var breweryCb breweryClipboard
	err := c.ShouldBindJSON(&breweryCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	brewery := models.Brewery{
		Name: breweryCb.Name,
	}

	err = uc.repository.UpdateBrewery(c, brewery)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "brewery update",
		"brewery": breweryCb,
	})
	return
}

func (uc *BreweryController) DeleteBrewery(c *gin.Context) {
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
	err = uc.repository.DeleteBrewery(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "brewery delete",
	})
}

//
