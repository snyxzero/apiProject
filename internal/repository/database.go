package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func New(c context.Context, connDbUrl string) (*Database, error) {
	pool, err := pgxpool.New(c, connDbUrl)
	if err != nil {
		return nil, fmt.Errorf("DB connection failed: %w", err)
	}
	return &Database{pool: pool}, nil
}

func (db *Database) Close() {
	db.pool.Close()
}

func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}
