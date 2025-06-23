package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snyxzero/apiProject/internal/models"
	"github.com/snyxzero/apiProject/internal/repository"
)

type userClipboard struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type UserController struct {
	repository *repository.UserRepository
}

func NewUserController(repository *repository.UserRepository) *UserController {
	return &UserController{
		repository: repository,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		// нужно возвращать в чем проблема, особено в ошибке запросе
		/*пример
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User ID must be greater than zero.", ну или на русском пофиг
		})*/
		return
	}
	if id < 1 {
		log.Println("incorrect id (id < 1)")
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := uc.repository.GetUser(c, id)
	if err != nil {
		log.Println(err)
		// нужно разделение через error.Is на типы ошибок
		// сейчас 2 типа
		// 1 ошибка работы с бд
		// 2 юзер не найден
		// интернал еррор только если ошибка с бд
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
	return
}

func (uc *UserController) CreateUser(c *gin.Context) {

	var userCb userClipboard
	err := c.ShouldBindJSON(&userCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	user := models.User{
		Name: userCb.Name,
	}

	userCb.ID, err = uc.repository.AddUser(c, user)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "user create",
		"user":   userCb,
	})
	return
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	var userCb userClipboard
	err := c.ShouldBindJSON(&userCb)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name: userCb.Name, // нет айдишника, обновляем по имени, это не верно, имя не уникальное
	}

	err = uc.repository.UpdateUser(c, user)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "user update",
		"user":   userCb,
	})
	return
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if id < 1 {
		log.Println("incorrect id (id < 1)")
		c.Status(http.StatusBadRequest)
		return
	}
	err = uc.repository.DeleteUser(c, id)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "user delete",
	})
}

//
