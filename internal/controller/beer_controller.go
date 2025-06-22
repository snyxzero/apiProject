package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
	"log"
	"net/http"
	"strconv"
)

type beerClipboard struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Brewery int    `json:"brewery" binding:"required"`
}

type BeerController struct {
	repository *repository.BeerRepository
}

func NewBeerController(repository *repository.BeerRepository) *BeerController {
	return &BeerController{
		repository: repository,
	}
}

func (o *BeerController) GetBeer(c *gin.Context) {
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
	beer, err := o.repository.GetBeer(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"beer":   beer,
	})
	return
}

func (o *BeerController) CreateBeer(c *gin.Context) {

	var beerCb beerClipboard
	err := c.ShouldBindJSON(&beerCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	beer := models.Beer{
		ID:      beerCb.ID,
		Name:    beerCb.Name,
		Brewery: beerCb.Brewery,
	}

	beerCb.ID, err = o.repository.AddBeer(c, beer)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "beer create",
		"beer":   beerCb,
	})
	return
}

func (o *BeerController) UpdateBeer(c *gin.Context) {

	var beerCb beerClipboard
	err := c.ShouldBindJSON(&beerCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	beer := models.Beer{
		ID:      beerCb.ID,
		Name:    beerCb.Name,
		Brewery: beerCb.Brewery,
	}

	err = o.repository.UpdateBeer(c, beer)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "beer update",
		"beer":   beerCb,
	})
	return
}

func (o *BeerController) DeleteBeer(c *gin.Context) {
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
	err = o.repository.DeleteBeer(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "beer delete",
	})
}
