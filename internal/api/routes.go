package api

import (
	"database/sql"
	"effectiveMobileTT/internal/repository"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	repo := repository.SongRepository{DB: db}
	handler := SongHandler{Repo: repo}

	e.GET("/songs", handler.GetSongs)
	e.GET("/songs/:id/text", handler.GetSongText)
	e.DELETE("/songs/:id", handler.DeleteSong)
	e.PUT("/songs/:id", handler.UpdateSong)
	e.POST("/songs", handler.AddSong)
}
