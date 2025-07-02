package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type UserBeerRatingService struct {
	userBeerRatingRepository *repository.UserBeerRatingsRepository
	userRepository           *repository.UsersRepository
	calculationRatingPoints  *СalculationRatingPoints
}

func NewUserBeerRatingService(userBeerRatingRepository *repository.UserBeerRatingsRepository, userRepository *repository.UsersRepository, calculationRatingPoints *СalculationRatingPoints) *UserBeerRatingService {
	return &UserBeerRatingService{
		userBeerRatingRepository: userBeerRatingRepository,
		userRepository:           userRepository,
		calculationRatingPoints:  calculationRatingPoints,
	}
}

func (o UserBeerRatingService) GetUserBeerRating(c *gin.Context, id int) (*models.UserBeerRating, error) {
	return o.userBeerRatingRepository.GetUserBeerRating(c, id)

}

func (o UserBeerRatingService) AddUserBeerRatingWithTransaction(c *gin.Context, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	// Начинаем транзакцию
	tx, err := o.userBeerRatingRepository.StartTransition(c)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	// Откладываем откат на случай ошибки
	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()
	// Выполняем операции в транзакции
	userBeerRating, err = o.userBeerRatingRepository.AddUserBeerRating(c, tx, userBeerRating)
	if err != nil {
		return nil, err
	}

	numberOfRatingsFromUser, err := o.userBeerRatingRepository.GetRatingCountForUser(c, tx, userBeerRating.ID)
	if err != nil {
		return nil, err
	}

	numberOfRatingsForBreweryFromUser, err := o.userBeerRatingRepository.GetRatingCountForUserForBrewery(c, tx, userBeerRating)
	if err != nil {
		return nil, err
	}

	points := o.calculationRatingPoints.CalculateRatingPointsToUser(numberOfRatingsFromUser, numberOfRatingsForBreweryFromUser)
	if err != nil {
		return nil, err
	}

	err = o.userRepository.UpdateUserPoints(c, tx, userBeerRating.User, points)
	if err != nil {
		return nil, err
	}
	// Фиксируем транзакцию
	if err := tx.Commit(c); err != nil {
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
