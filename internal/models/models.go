package models

// Song представляет информацию о песне.
// @Description Структура для отображения основной информации о песне.
type Song struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Name        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
}

// NewSong представляет полную информацию о песне, включая текст и ссылку.
// @Description Структура для отображения полной информации о песне.
type NewSong struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// UpdateSongRequest содержит поля для обновления информации о песне.
// @Description Структура для обновления данных о песне.
type UpdateSongRequest struct {
	Group string `json:"group,omitempty"`
	Song  string `json:"song,omitempty"`
}

// AddSongRequest содержит данные для добавления новой песни.
// @Description Структура для добавления новой песни.
// @Param group body string true "Название группы"
// @Param song body string true "Название песни"
type AddSongRequest struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}

// ExternalAPIResponse представляет ответ внешнего API.
// @Description Структура ответа от внешнего API с дополнительной информацией о песне.
type ExternalAPIResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// DBConnection содержит параметры подключения к базе данных.
// @Description Структура для хранения конфигурации подключения к базе данных.
type DBConnection struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}
