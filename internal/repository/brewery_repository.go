package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/models"
)

type BreweryRepository struct {
	pool *pgxpool.Pool
}

func NewBreweryRepository(pool *pgxpool.Pool) *BreweryRepository {
	return &BreweryRepository{pool: pool}
}

func (o *BreweryRepository) GetBrewery(ctx context.Context, id int) (*models.Brewery, error) {
	var brewery models.Brewery
	err := o.pool.QueryRow(ctx, "SELECT id, name FROM breweries WHERE id = $1", id).Scan(&brewery.ID, &brewery.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get brewery id %d from db: %v ", id, err)
	}
	return &brewery, nil
}

func (o *BreweryRepository) AddBrewery(ctx context.Context, brewery models.Brewery) (int, error) {
	err := o.pool.QueryRow(ctx, "INSERT INTO breweries (name) VALUES ($1) RETURNING id", brewery.Name).Scan(&brewery.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to add brewery to db: %v", err)
	}
	return brewery.ID, nil
}

func (o *BreweryRepository) UpdateBrewery(ctx context.Context, brewery models.Brewery) error {
	_, err := o.pool.Exec(ctx, "UPDATE breweries SET name = $1 WHERE id = $2", brewery.Name, brewery.ID)
	if err != nil {
		return fmt.Errorf("failed to update brewery in db: %v", err)
	}
	return nil
}

func (o *BreweryRepository) DeleteBrewery(ctx context.Context, id int) error {
	_, err := o.pool.Exec(ctx, `DELETE FROM breweries WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete brewery in db: %v", err)
	}
	return nil
}

//
