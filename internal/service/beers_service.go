package service

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type BeerService struct {
	repository *repository.BeersRepository
}

func NewBeerService(repository *repository.BeersRepository) *BeerService {
	return &BeerService{
		repository: repository,
	}
}

func (o BeerService) GetBeer(c *gin.Context, id int) (*models.Beer, error) {
	return o.repository.GetBeer(c, id)
}

func (o BeerService) AddBeer(c *gin.Context, brewery *models.Beer) (*models.Beer, error) {
	return o.repository.AddBeer(c, brewery)
}

func (o BeerService) UpdateBeer(c *gin.Context, brewery *models.Beer) (*models.Beer, error) {
	return o.repository.UpdateBeer(c, brewery)
}

func (o BeerService) DeleteBeer(c *gin.Context, id int) error {
	return o.repository.DeleteBeer(c, id)
}
