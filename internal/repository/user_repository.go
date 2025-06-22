package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snyxzero/apiProject/internal/models"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (o *UserRepository) GetUser(c *gin.Context, id int) (*models.User, error) {
	var user models.User
	err := o.pool.QueryRow(c, "SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id %d from db: %v ", id, err)
	}
	return &user, nil
}

func (o *UserRepository) AddUser(c *gin.Context, user models.User) (int, error) {
	err := o.pool.QueryRow(c, "INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to add user to db: %v", err)
	}
	return user.ID, nil
}

func (o *UserRepository) UpdateUser(c *gin.Context, user models.User) error {
	_, err := o.pool.Exec(c, "UPDATE users SET name = $1 WHERE id = $2", user.Name, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user in db: %v", err)
	}
	return nil
}

func (o *UserRepository) DeleteUser(c *gin.Context, id int) error {
	_, err := o.pool.Exec(c, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user in db: %v", err)
	}
	return nil
}
