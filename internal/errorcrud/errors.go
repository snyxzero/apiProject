package errorcrud

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	ErrInvalidFormat          = errors.New("invalid id format")
	ErrNegativeID             = errors.New("id must be positive")
	ErrUserNotFound           = errors.New("user not found")
	ErrBeerNotFound           = errors.New("beer not found")
	ErrBreweryNotFound        = errors.New("brewery not found")
	ErrUserBeerRatingNotFound = errors.New("rating not found")
	ErrCreatingData           = errors.New("error creating data in db")
	ErrGettingData            = errors.New("error getting data in db")
	ErrUpdatingData           = errors.New("error updating data in db")
	ErrDeletingData           = errors.New("error deleting data in db")
)

func ErrorCheck(c *gin.Context, err error) {
	ErrUpdatingData.Error()
	switch {
	case errors.Is(err, ErrInvalidFormat):
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid ID format",
			"details": err.Error(),
		})

	case errors.Is(err, ErrNegativeID):
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID must be positive",
			"details": err.Error(),
		})

	case errors.Is(err, ErrUserNotFound):
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "User not found",
			"details": err.Error(),
		})

	case errors.Is(err, ErrBeerNotFound):
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Beer not found",
			"details": err.Error(),
		})

	case errors.Is(err, ErrBreweryNotFound):
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Brewery not found",
			"details": err.Error(),
		})

	case errors.Is(err, ErrUserBeerRatingNotFound):
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "UserBeerRating not found",
			"details": err.Error(),
		})

	case errors.Is(err, ErrGettingData):
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error getting data in db",
			"details": err.Error(),
		})

	case errors.Is(err, ErrCreatingData):
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating data in db",
			"details": err.Error(),
		})

	case errors.Is(err, ErrUpdatingData):
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating data in db",
			"details": err.Error(),
		})

	case errors.Is(err, ErrDeletingData):
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error deleting data in db",
			"details": err.Error(),
		})

	default:
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Unknown error",
			"details": err.Error(),
		})
	}
	return
}

func ErrInvalidJson(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   "Invalid json data ",
		"details": err.Error(),
	})
	return
}
