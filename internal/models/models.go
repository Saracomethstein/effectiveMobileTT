package models

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Name        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
}

type NewSong struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type UpdateSongRequest struct {
	Group string `json:"group,omitempty"`
	Song  string `json:"song,omitempty"`
}

type AddSongRequest struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}

type ExternalAPIResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type DBConnection struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}
