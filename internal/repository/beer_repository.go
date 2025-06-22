package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/models"
)

type BeerRepository struct {
	pool *pgxpool.Pool
}

func NewBeerRepository(pool *pgxpool.Pool) *BeerRepository {
	return &BeerRepository{pool: pool}
}

func (o *BeerRepository) GetBeer(ctx context.Context, id int) (*models.Beer, error) {
	var beer models.Beer
	err := o.pool.QueryRow(ctx, `
SELECT id, name, brewery_id 
FROM beers_types 
WHERE id = $1`, id).Scan(&beer.ID, &beer.Name, &beer.Brewery)
	if err != nil {
		return nil, fmt.Errorf("failed to get beer id %d from db: %v ", id, err)
	}
	return &beer, nil
}

func (o *BeerRepository) AddBeer(ctx context.Context, beer models.Beer) (int, error) {
	err := o.pool.QueryRow(ctx, `
INSERT INTO beers_types (name, brewery_id) 
VALUES ($1, $2) RETURNING id`, beer.Name, beer.Brewery).Scan(&beer.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to add beer to db: %v", err)
	}
	return beer.ID, nil
}

func (o *BeerRepository) UpdateBeer(ctx context.Context, beer models.Beer) error {
	_, err := o.pool.Exec(ctx, `
UPDATE beers_types 
SET name = $1, brewery_id = $2 WHERE id = $3`, beer.Name, beer.Brewery, beer.ID)
	if err != nil {
		return fmt.Errorf("failed to update beer in db: %v", err)
	}
	return nil
}

func (o *BeerRepository) DeleteBeer(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `
DELETE FROM beers_types 
WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete beer in db: %v", err)
	}
	return nil
}

//
