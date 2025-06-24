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

type BeersRepository struct {
	pool *pgxpool.Pool
}

func NewBeersRepository(pool *pgxpool.Pool) *BeersRepository {
	return &BeersRepository{pool: pool}
}

func (o *BeersRepository) GetBeer(ctx context.Context, id int) (*models.Beer, error) {
	var beer models.Beer
	err := o.pool.QueryRow(ctx, `
SELECT id, name, breweries_id FROM beers 
WHERE id = $1`, id).Scan(&beer.ID, &beer.Name, &beer.Brewery)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorcrud.ErrBeerNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrGettingData, err)
	}
	return &beer, nil
}

func (o *BeersRepository) AddBeer(ctx context.Context, beer *models.Beer) (models.Beer, error) {
	err := o.pool.QueryRow(ctx, `
INSERT INTO beers (name, breweries_id) 
VALUES ($1, $2) RETURNING id, name, breweries_id`, beer.Name, beer.Brewery).Scan(&beer.ID, &beer.Name, &beer.Brewery)
	if err != nil {
		return models.Beer{}, fmt.Errorf("%w: %v", errorcrud.ErrCreatingData, err)
	}
	return *beer, nil
}

func (o *BeersRepository) UpdateBeer(ctx context.Context, beer *models.Beer) (models.Beer, error) {
	err := o.pool.QueryRow(ctx, `
UPDATE beers SET name = $1, breweries_id = $2 WHERE id = $3 
RETURNING id, name, breweries_id`, beer.Name, beer.Brewery, beer.ID).Scan(&beer.ID, &beer.Name, &beer.Brewery)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Beer{}, fmt.Errorf("%w: %v", errorcrud.ErrBeerNotFound, err)
		}
		return models.Beer{}, fmt.Errorf("%w: %v", errorcrud.ErrUpdatingData, err)
	}
	return *beer, nil
}

func (o *BeersRepository) DeleteBeer(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `
DELETE FROM beers 
WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%w: %v", errorcrud.ErrBeerNotFound, err)
		}
		return fmt.Errorf("%w: %v", errorcrud.ErrDeletingData, err)
	}
	return nil
}

//
