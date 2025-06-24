package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snyxzero/apiProject/internal/controller"
	"github.com/snyxzero/apiProject/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	// Подключение к БД
	db, err := repository.New(ctx, "postgres://myuser2:123@localhost:5432/mydb2")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usersRepo := repository.NewUsersRepository(db.Pool())
	usersCtrl := controller.NewUserController(usersRepo)
	// Группа маршрутов /api/users
	apiUsers := r.Group("/api/users")
	{
		apiUsers.POST("/", usersCtrl.CreateUser)
		apiUsers.GET("/:id", usersCtrl.GetUser)
		apiUsers.PUT("/:id", usersCtrl.UpdateUser)
		apiUsers.DELETE("/:id", usersCtrl.DeleteUser)
	}

	breweriesRepo := repository.NewBreweriesRepository(db.Pool())
	breweriesCtrl := controller.NewBreweryController(breweriesRepo)
	// Группа маршрутов /api/breweries
	apiBreweries := r.Group("/api/breweries")
	{
		apiBreweries.POST("/", breweriesCtrl.CreateBrewery)
		apiBreweries.GET("/:id", breweriesCtrl.GetBrewery)
		apiBreweries.PUT("/:id", breweriesCtrl.UpdateBrewery)
		apiBreweries.DELETE("/:id", breweriesCtrl.DeleteBrewery)
	}

	beersRepo := repository.NewBeersRepository(db.Pool())
	beersCtrl := controller.NewBeerController(beersRepo)
	// Группа маршрутов /api/beers
	apiBeers := r.Group("/api/beers")
	{
		apiBeers.POST("/", beersCtrl.CreateBeer)
		apiBeers.GET("/:id", beersCtrl.GetBeer)
		apiBeers.PUT("/:id", beersCtrl.UpdateBeer)
		apiBeers.DELETE("/:id", beersCtrl.DeleteBeer)
	}

	usersBeersRatingRepo := repository.NewUserBeerRatingsRepository(db.Pool())
	usersBeersRatingCtrl := controller.NewRatingController(usersBeersRatingRepo)
	// Группа маршрутов /api/usersbeersrating
	apiUsersBeersRating := r.Group("/api/usersbeersrating")
	{
		apiUsersBeersRating.POST("/", usersBeersRatingCtrl.CreateRating)
		apiUsersBeersRating.GET("/:id", usersBeersRatingCtrl.GetRating)
		apiUsersBeersRating.PUT("/:id", usersBeersRatingCtrl.UpdateRating)
		apiUsersBeersRating.DELETE("/:id", usersBeersRatingCtrl.DeleteRating)
	}

	// Создание сервера
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		fmt.Println("Сервер запущен на http://localhost:8080")
		log.Fatal(srv.ListenAndServe())
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала завершения
	<-stop

	// Graceful Shutdown

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Forced shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}

//
