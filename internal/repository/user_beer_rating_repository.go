package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/models"
)

type UserBeerRatingRepository struct {
	pool *pgxpool.Pool
}

func NewRatingRepository(pool *pgxpool.Pool) *UserBeerRatingRepository {
	return &UserBeerRatingRepository{pool: pool}
}

func (o *UserBeerRatingRepository) GetRating(ctx context.Context, id int) (*models.UserBeerRating, error) {
	var rating models.UserBeerRating
	err := o.pool.QueryRow(ctx, `
SELECT id, user_id, beer_id, rating FROM user_beer_ratings 
WHERE id = $1`, id).Scan(&rating.ID, &rating.User, &rating.Beer, &rating.Rating)
	if err != nil {
		return nil, fmt.Errorf("failed to get user_beer_ratings id %d from db: %v ", id, err)
	}
	return &rating, nil
}

func (o *UserBeerRatingRepository) AddRating(ctx context.Context, rating models.UserBeerRating) (int, error) {
	err := o.pool.QueryRow(ctx, `
INSERT INTO user_beer_ratings (user_id, beer_id, rating) 
VALUES ($1, $2, $3) RETURNING id`, rating.User, rating.Beer, rating.Rating).Scan(&rating.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to add user_beer_ratings to db: %v", err)
	}
	return rating.ID, nil
}

func (o *UserBeerRatingRepository) UpdateRating(ctx context.Context, rating models.UserBeerRating) error {
	_, err := o.pool.Exec(ctx, `
UPDATE user_beer_ratings 
SET user_id = $2, beer_id = $3, rating = $4 WHERE id = $1`, rating.ID, rating.User, rating.Beer, rating.Rating)
	if err != nil {
		return fmt.Errorf("failed to update user_beer_ratings in db: %v", err)
	}
	return nil
}

func (o *UserBeerRatingRepository) DeleteRating(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `
DELETE FROM user_beer_ratings 
WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user_beer_ratings in db: %v", err)
	}
	return nil
}
