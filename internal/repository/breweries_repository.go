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

type BreweriesRepository struct {
	pool *pgxpool.Pool
}

func NewBreweriesRepository(pool *pgxpool.Pool) *BreweriesRepository {
	return &BreweriesRepository{pool: pool}
}

func (o *BreweriesRepository) GetBrewery(ctx context.Context, id int) (*models.Brewery, error) {
	var brewery models.Brewery
	err := o.pool.QueryRow(ctx, `
SELECT id, name FROM breweries 
WHERE id = $1`, id).Scan(&brewery.ID, &brewery.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorcrud.ErrBreweryNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrGettingData, err)
	}
	return &brewery, nil
}

func (o *BreweriesRepository) AddBrewery(ctx context.Context, brewery *models.Brewery) (*models.Brewery, error) {
	err := o.pool.QueryRow(ctx, `
INSERT INTO breweries (name) VALUES ($1) 
RETURNING id, name`, brewery.Name).Scan(&brewery.ID, &brewery.Name)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrCreatingData, err)
	}
	return brewery, nil
}

func (o *BreweriesRepository) UpdateBrewery(ctx context.Context, brewery *models.Brewery) (*models.Brewery, error) {
	err := o.pool.QueryRow(ctx, `
UPDATE breweries SET name = $1 WHERE id = $2 
RETURNING id, name`, brewery.Name, brewery.ID).Scan(&brewery.ID, &brewery.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorcrud.ErrBreweryNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrUpdatingData, err)
	}
	return brewery, nil
}

func (o *BreweriesRepository) DeleteBrewery(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `DELETE FROM breweries WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%w: %v", errorcrud.ErrBreweryNotFound, err)
		}
		return fmt.Errorf("%w: %v", errorcrud.ErrDeletingData, err)
	}
	return nil
}

//
