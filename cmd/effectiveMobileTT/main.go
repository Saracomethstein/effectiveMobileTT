package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"

	_ "effectiveMobileTT/docs"
	"effectiveMobileTT/internal/api"
	"effectiveMobileTT/internal/repository"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @title Songs API
// @version 1.0
// @description API для управления песнями, включая добавление, обновление, удаление и получение списка песен.
// @host localhost:8000
// @BasePath /
func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := repository.SetupDB()
	defer db.Close()

	repo := repository.SongRepository{DB: db}
	handler := api.SongHandler{Repo: repo}

	e.GET("/songs", handler.GetSongs)
	e.GET("/songs/:id/text", handler.GetSongText)
	e.DELETE("/songs/:id", handler.DeleteSong)
	e.PUT("/songs/:id", handler.UpdateSong)
	e.POST("/songs", handler.AddSong)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Println("Server is running on port 8000")
	log.Fatal(e.Start(":8000"))
}
