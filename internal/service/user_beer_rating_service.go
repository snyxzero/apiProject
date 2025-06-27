package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type UserBeerRatingService struct {
	userBeerRatingRepository *repository.UserBeerRatingsRepository
	ratingPoints             *RatingPoints
}

func NewUserBeerRatingService(repositoryUserBeerRating *repository.UserBeerRatingsRepository, ratingPoints *RatingPoints) *UserBeerRatingService {
	return &UserBeerRatingService{
		userBeerRatingRepository: repositoryUserBeerRating,
		ratingPoints:             ratingPoints,
	}
}

func (o UserBeerRatingService) GetUserBeerRating(c *gin.Context, id int) (*models.UserBeerRating, error) {
	return o.userBeerRatingRepository.GetUserBeerRating(c, id)

}

func (o UserBeerRatingService) AddUserBeerRatingWithTransaction(c *gin.Context, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	ctx := c.Request.Context()

	// Начинаем транзакцию
	tx, err := o.userBeerRatingRepository.StartTransition(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Откладываем откат на случай ошибки
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Выполняем операции в транзакции
	userBeerRating, err = o.userBeerRatingRepository.AddUserBeerRating(c, tx, userBeerRating)
	if err != nil {
		return nil, err
	}

	err = o.ratingPoints.AddRatingPointsToUser(c, tx, userBeerRating)
	if err != nil {
		return nil, err
	}

	// Фиксируем транзакцию
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return userBeerRating, nil
}

func (o UserBeerRatingService) UpdateUserBeerRating(c *gin.Context, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	return o.userBeerRatingRepository.UpdateUserBeerRating(c, userBeerRating)
}

func (o UserBeerRatingService) DeleteUserBeerRating(c *gin.Context, id int) error {
	return o.userBeerRatingRepository.DeleteUserBeerRating(c, id)
}
