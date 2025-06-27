package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type UserBeerRatingService struct {
	repository *repository.UserBeerRatingsRepository
}

func NewUserBeerRatingService(repository *repository.UserBeerRatingsRepository) *UserBeerRatingService {
	return &UserBeerRatingService{
		repository: repository,
	}
}

func (o UserBeerRatingService) GetUserBeerRating(c *gin.Context, id int) (*models.UserBeerRating, error) {
	return o.repository.GetUserBeerRating(c, id)
}

func (o UserBeerRatingService) AddUserBeerRating(c *gin.Context, tx pgx.Tx, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	return o.repository.AddUserBeerRating(c, tx, userBeerRating)
}

func (o UserBeerRatingService) UpdateUserBeerRating(c *gin.Context, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	return o.repository.UpdateUserBeerRating(c, userBeerRating)
}

func (o UserBeerRatingService) DeleteUserBeerRating(c *gin.Context, id int) error {
	return o.repository.DeleteUserBeerRating(c, id)
}
