package userbeerrating

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type RatingPoints struct {
	ratingRepository repository.UserBeerRatingsRepository
	userRepository   repository.UsersRepository
}

func NewRatingPoints(ratingRepository *repository.UserBeerRatingsRepository, userRepository *repository.UsersRepository) *RatingPoints {
	return &RatingPoints{
		ratingRepository: *ratingRepository,
		userRepository:   *userRepository,
	}
}

func (r *RatingPoints) AddRatingPointsToUser(ctx *gin.Context, tx pgx.Tx, userBeerRating *models.UserBeerRating) error {
	countRating, err := r.ratingRepository.GetRatingCountForUser(ctx, tx, userBeerRating.ID)
	if err != nil {
		return err
	}

	countRatingForBrewery, err := r.ratingRepository.GetRatingCountForUserForBrewery(ctx, tx, userBeerRating)
	if err != nil {
		return err
	}

	points := 0

	if countRating == 1 {
		points += 50
	}

	if countRating%3 == 0 {
		points += 5
	}

	if countRatingForBrewery%2 == 0 {
		points += 10
	}
	err = r.userRepository.UpdateUserPoints(ctx, tx, userBeerRating.ID, points)
	return nil
}

func (o *RatingPoints) AddRatingWithTransaction(c *gin.Context, userBeerRating *models.UserBeerRating) (*models.UserBeerRating, error) {
	ctx := c.Request.Context()

	// Начинаем транзакцию
	tx, err := o.ratingRepository.StartTransition(ctx)
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
	result, err := o.ratingRepository.AddUserBeerRating(c, tx, userBeerRating)
	if err != nil {
		return nil, err
	}

	err = o.AddRatingPointsToUser(c, tx, result)
	if err != nil {
		return nil, err
	}

	// Фиксируем транзакцию
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return result, nil
}
