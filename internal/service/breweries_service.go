package service

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type BreweryService struct {
	repository *repository.BreweriesRepository
}

func NewBreweryService(repository *repository.BreweriesRepository) *BreweryService {
	return &BreweryService{
		repository: repository,
	}
}

func (o BreweryService) GetBrewery(c *gin.Context, id int) (*models.Brewery, error) {
	return o.repository.GetBrewery(c, id)
}

func (o BreweryService) AddBrewery(c *gin.Context, brewery *models.Brewery) (*models.Brewery, error) {
	return o.repository.AddBrewery(c, brewery)
}

func (o BreweryService) UpdateBrewery(c *gin.Context, brewery *models.Brewery) (*models.Brewery, error) {
	return o.repository.UpdateBrewery(c, brewery)
}

func (o BreweryService) DeleteBrewery(c *gin.Context, id int) error {
	return o.repository.DeleteBrewery(c, id)
}
