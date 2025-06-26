package repository

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/errorcrud"
	"github.com/snyxzero/apiProject/internal/models"
)

type UsersRepository struct {
	pool *pgxpool.Pool
}

func NewUsersRepository(pool *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{pool: pool}
}

func (o *UsersRepository) GetUser(c *gin.Context, id int) (*models.User, error) {
	var user models.User
	err := o.pool.QueryRow(c, "SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.RatingPoints)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorcrud.ErrUserNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", errorcrud.ErrGettingData, err)
	}
	return &user, nil
}

func (o *UsersRepository) AddUser(c *gin.Context, user *models.User) (models.User, error) {
	err := o.pool.QueryRow(c, `INSERT INTO users (name, rating_points) VALUES ($1, 0) RETURNING *`, user.Name).Scan(&user.ID, &user.Name, &user.RatingPoints)
	if err != nil {
		return models.User{}, fmt.Errorf("%w: %v", errorcrud.ErrCreatingData, err)
	}
	return *user, nil
}

func (o *UsersRepository) UpdateUser(c *gin.Context, user *models.User) (models.User, error) {
	err := o.pool.QueryRow(c, `
UPDATE users SET name = $1 WHERE id = $2 
RETURNING *`, user.Name, user.ID).Scan(&user.ID, &user.Name, &user.RatingPoints)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("%w: %v", errorcrud.ErrUserNotFound, err)
		}
		return models.User{}, fmt.Errorf("%w: %v", errorcrud.ErrUpdatingData, err)
	}
	return *user, nil
}

func (o *UsersRepository) DeleteUser(c *gin.Context, id int) error {
	_, err := o.pool.Exec(c, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%w: %v", errorcrud.ErrUserNotFound, err)
		}
		return fmt.Errorf("%w: %v", errorcrud.ErrDeletingData, err)
	}
	return nil
}

func (o *UsersRepository) UpdateUserPoints(c *gin.Context, tx pgx.Tx, id int, points int) error {
	if points == 0 {
		return nil
	}

	_, err := tx.Exec(c, `
UPDATE users SET rating_points = $2 WHERE id = $1 
RETURNING *`, id, points)
	if err != nil {
		return fmt.Errorf("%w: %v", errorcrud.ErrUpdatingData, err)
	}
	return nil
}

//
