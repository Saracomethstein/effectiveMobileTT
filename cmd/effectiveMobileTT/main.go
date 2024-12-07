package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"

	"effectiveMobileTT/internal/api"
	"effectiveMobileTT/internal/repository"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := repository.SetupDB()
	defer db.Close()

	api.InitRoutes(e, db)

	// Запуск сервера
	log.Println("Server is running on port 8000")
	log.Fatal(e.Start(":8000"))
}
