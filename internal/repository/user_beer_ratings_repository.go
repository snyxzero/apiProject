package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/models"
)

type UserBeerRatingsRepository struct {
	pool *pgxpool.Pool
}

func NewUserBeerRatingsRepository(pool *pgxpool.Pool) *UserBeerRatingsRepository {
	return &UserBeerRatingsRepository{pool: pool}
}

func (o *UserBeerRatingsRepository) GetRating(ctx context.Context, id int) (*models.UserBeerRating, error) {
	var rating models.UserBeerRating
	err := o.pool.QueryRow(ctx, `
SELECT id, users_id, beers_id, rating FROM user_beer_ratings 
WHERE id = $1`, id).Scan(&rating.ID, &rating.User, &rating.Beer, &rating.Rating)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorcrud.ErrUserBeerRatingNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrGettingData, err)
	}
	return &rating, nil
}

func (o *UserBeerRatingsRepository) AddRating(ctx context.Context, rating *models.UserBeerRating) (models.UserBeerRating, error) {
	err := o.pool.QueryRow(ctx, `
INSERT INTO user_beer_ratings (users_id, beers_id, rating) 
VALUES ($1, $2, $3) RETURNING id, users_id, beers_id, rating`, rating.User, rating.Beer, rating.Rating).Scan(&rating.ID, &rating.User, &rating.Beer, &rating.Rating)
	if err != nil {
		return models.UserBeerRating{}, fmt.Errorf("%w: %v", errorcrud.ErrCreatingData, err)
	}
	return *rating, nil
}

func (o *UserBeerRatingsRepository) UpdateRating(ctx context.Context, rating *models.UserBeerRating) (models.UserBeerRating, error) {
	err := o.pool.QueryRow(ctx, `
UPDATE user_beer_ratings 
SET users_id = $2, beers_id = $3, rating = $4 
WHERE id = $1
RETURNING id, users_id, beers_id, rating`, rating.ID, rating.User, rating.Beer, rating.Rating).Scan(&rating.ID, &rating.User, &rating.Beer, &rating.Rating)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserBeerRating{}, fmt.Errorf("%w: %v", errorcrud.ErrUserBeerRatingNotFound, err)
		}
		return models.UserBeerRating{}, fmt.Errorf("%w: %v", errorcrud.ErrUpdatingData, err)
	}
	return *rating, nil
}

func (o *UserBeerRatingsRepository) DeleteRating(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `
DELETE FROM user_beer_ratings 
WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%w: %v", errorcrud.ErrUserBeerRatingNotFound, err)
		}
		return fmt.Errorf("%w: %v", errorcrud.ErrDeletingData, err)
	}
	return nil
}

//
