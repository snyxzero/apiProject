package ratingpoints

import (
	"github.com/gin-gonic/gin"
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

func (r *RatingPoints) AddRatingPointsToUser(ctx *gin.Context, userBeerRating *models.UserBeerRating) error {
	countRating, err := r.ratingRepository.GetRatingCountForUser(ctx, userBeerRating.ID)
	if err != nil {
		return err
	}

	countRatingForBrewery, err := r.ratingRepository.GetRatingCountForUserForBrewery(ctx, userBeerRating)
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
	err = r.userRepository.UpdateUserPoints(ctx, userBeerRating.ID, points)
	return nil
}
