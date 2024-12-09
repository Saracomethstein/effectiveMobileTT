package api

import (
	_ "effectiveMobileTT/cmd/effectiveMobileTT/docs"
	"effectiveMobileTT/internal/models"
	"effectiveMobileTT/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SongHandler struct {
	Repo repository.SongRepository
}

// GetSongs получает список песен с фильтрацией и пагинацией.
// @Summary Get songs
// @Description Получение списка песен с фильтрацией по группе, названию и дате релиза.
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Название группы"
// @Param song query string false "Название песни"
// @Param releaseDate query string false "Дата релиза (в формате YYYY-MM-DD)"
// @Param limit query int false "Количество записей (по умолчанию 10)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {array} models.Song
// @Failure 500 {object} echo.Map "Ошибка сервера"
// @Router /songs [get]
func (h *SongHandler) GetSongs(c echo.Context) error {
	group := c.QueryParam("group")
	song := c.QueryParam("song")
	releaseDate := c.QueryParam("releaseDate")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := 10
	offset := 0

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	songs, err := h.Repo.GetSongs(group, song, releaseDate, limit, offset)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch songs",
		})
	}

	return c.JSON(http.StatusOK, songs)
}

// GetSongText получает текст песни с пагинацией по куплетам.
// @Summary Get song text
// @Description Получение текста песни по ID с возможностью пагинации по куплетам.
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "ID песни"
// @Param limit query int false "Количество куплетов (по умолчанию 5)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {array} string "Список куплетов"
// @Failure 404 {object} echo.Map "Песня не найдена"
// @Router /songs/{id}/text [get]
func (h *SongHandler) GetSongText(c echo.Context) error {
	id := c.Param("id")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := 5
	offset := 0

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	text, err := h.Repo.GetSongTextByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Song not found",
		})
	}

	verses := strings.Split(text, "\n\n")
	if offset >= len(verses) {
		return c.JSON(http.StatusOK, []string{})
	}

	end := offset + limit
	if end > len(verses) {
		end = len(verses)
	}

	paginatedVerses := verses[offset:end]
	return c.JSON(http.StatusOK, paginatedVerses)
}

// DeleteSong удаляет песню по ID.
// @Summary Delete a song
// @Description Удаление песни по её ID.
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "ID песни"
// @Success 204 "Песня удалена"
// @Failure 404 {object} echo.Map "Песня не найдена"
// @Failure 500 {object} echo.Map "Ошибка сервера"
// @Router /songs/{id} [delete]
func (h *SongHandler) DeleteSong(c echo.Context) error {
	id := c.Param("id")

	err := h.Repo.DeleteSongByID(id)
	if err != nil {
		if err.Error() == "song not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Song not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to delete song",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// UpdateSong обновляет информацию о песне по ID.
// @Summary Update a song
// @Description Обновление информации о песне по её ID.
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "ID песни"
// @Param body body models.UpdateSongRequest true "Данные для обновления"
// @Success 200 {object} models.Song "Обновлённая песня"
// @Failure 400 {object} echo.Map "Неверный запрос"
// @Failure 404 {object} echo.Map "Песня не найдена"
// @Failure 500 {object} echo.Map "Ошибка сервера"
// @Router /songs/{id} [put]
func (h *SongHandler) UpdateSong(c echo.Context) error {
	id := c.Param("id")

	var req models.UpdateSongRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	updatedSong, err := h.Repo.UpdateSongByID(id, req.Group, req.Song)
	if err != nil {
		if err.Error() == "song not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Song not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to update song",
		})
	}

	return c.JSON(http.StatusOK, updatedSong)
}

// AddSong добавляет новую песню в базу данных.
// @Summary Add a song
// @Description Добавление новой песни с получением данных из внешнего API.
// @Tags Songs
// @Accept json
// @Produce json
// @Param body body models.AddSongRequest true "Данные для добавления песни"
// @Success 201 {object} models.Song "Созданная песня"
// @Failure 400 {object} echo.Map "Неверный запрос или ошибка валидации"
// @Failure 500 {object} echo.Map "Ошибка сервера или внешнего API"
// @Router /songs [post]
func (h *SongHandler) AddSong(c echo.Context) error {
	var req models.AddSongRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Validation failed",
		})
	}

	apiURL := fmt.Sprintf("http://external-api-url/info?group=%s&song=%s", req.Group, req.Song)
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch song details from external API",
		})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to read external API response",
		})
	}

	var apiData models.ExternalAPIResponse
	if err := json.Unmarshal(body, &apiData); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Invalid response from external API",
		})
	}

	newSong, err := h.Repo.AddSong(req.Group, req.Song, apiData.ReleaseDate, apiData.Text, apiData.Link)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to save song to database",
		})
	}

	return c.JSON(http.StatusCreated, newSong)
}
