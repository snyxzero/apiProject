package service

import (
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type UserService struct {
	repository *repository.UsersRepository
}

func NewUserService(repository *repository.UsersRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (o UserService) GetUser(c *gin.Context, id int) (*models.User, error) {
	return o.repository.GetUser(c, id)
}

func (o UserService) AddUser(c *gin.Context, user *models.User) (*models.User, error) {
	return o.repository.AddUser(c, user)
}

func (o UserService) UpdateUser(c *gin.Context, user *models.User) (*models.User, error) {
	return o.repository.UpdateUser(c, user)
}

func (o UserService) DeleteUser(c *gin.Context, id int) error {
	return o.repository.DeleteUser(c, id)
}
