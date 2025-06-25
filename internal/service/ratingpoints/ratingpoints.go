package ratingpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type RatingPoints struct {
	repositoryRating repository.UserBeerRatingsRepository
	repositoryUser   repository.UsersRepository
}

func NewRatingPoints(repositoryRating *repository.UserBeerRatingsRepository, repositoryUser *repository.UsersRepository) *RatingPoints {
	return &RatingPoints{
		repositoryRating: *repositoryRating,
		repositoryUser:   *repositoryUser,
	}
}

func (r *RatingPoints) AddRatingPointsToUser(ctx *gin.Context, userBeerRating *models.UserBeerRating) error {
	countRating, err := r.repositoryRating.GetRatingCountForUser(ctx, userBeerRating.ID)
	if err != nil {
		return err
	}

	countRatingForBrewery, err := r.repositoryRating.GetRatingCountForUserForBrewery(ctx, userBeerRating)
	if err != nil {
		return err
	}

	if countRating == 1 {
		err = r.repositoryUser.UpdateUserPoints(ctx, userBeerRating.ID, 50)
		if err != nil {
			return err
		}
	}

	if countRating%3 == 0 {
		err = r.repositoryUser.UpdateUserPoints(ctx, userBeerRating.ID, 5)
		if err != nil {
			return err
		}
	}

	if countRatingForBrewery%2 == 0 {
		err = r.repositoryUser.UpdateUserPoints(ctx, userBeerRating.ID, 10)
		if err != nil {
			return err
		}
	}
	return nil
}
