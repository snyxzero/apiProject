package service

import (
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
	numberOfRatingsFromUser, err := r.ratingRepository.GetRatingCountForUser(ctx, tx, userBeerRating.ID)
	if err != nil {
		return err
	}

	numberOfRatingsForBreweryFromUser, err := r.ratingRepository.GetRatingCountForUserForBrewery(ctx, tx, userBeerRating)
	if err != nil {
		return err
	}

	points := 0

	if numberOfRatingsFromUser == 1 {
		points += 50
	}

	if numberOfRatingsFromUser%3 == 0 {
		points += 5
	}

	if numberOfRatingsForBreweryFromUser%2 == 0 {
		points += 10
	}
	err = r.userRepository.UpdateUserPoints(ctx, tx, userBeerRating.ID, points)
	return nil
}
