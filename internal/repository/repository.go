package repository

import (
	"context"
	"database/sql"
	"effectiveMobileTT/internal/models"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type SongRepository struct {
	DB *sql.DB
}

const (
	retries = 5
	delay   = 3 * time.Second
)

func (r *SongRepository) GetSongs(group, song, releaseDate string, limit, offset int) ([]models.Song, error) {
	query := "SELECT id, group_name, song_name, release_date FROM songs WHERE 1=1"
	args := []interface{}{}

	if group != "" {
		query += " AND group_name ILIKE ?"
		args = append(args, "%"+group+"%")
	}
	if song != "" {
		query += " AND song_name ILIKE ?"
		args = append(args, "%"+song+"%")
	}
	if releaseDate != "" {
		query += " AND release_date = ?"
		args = append(args, releaseDate)
	}

	query += " LIMIT $1 OFFSET $2"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Group, &song.Name, &song.ReleaseDate)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongRepository) GetSongTextByID(id string) (string, error) {
	query := "SELECT text FROM songs WHERE id = $1"

	var text string
	err := r.DB.QueryRowContext(context.Background(), query, id).Scan(&text)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return text, nil
}

func (r *SongRepository) DeleteSongByID(id string) error {
	queryCheck := "SELECT COUNT(*) FROM songs WHERE id = $1"
	var count int
	err := r.DB.QueryRowContext(context.Background(), queryCheck, id).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("song not found")
	}

	queryDelete := "DELETE FROM songs WHERE id = $1"
	_, err = r.DB.ExecContext(context.Background(), queryDelete, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SongRepository) UpdateSongByID(id, group, song string) (*models.Song, error) {
	queryCheck := "SELECT id, group_name, song_name FROM songs WHERE id = $1"
	var existingSong models.Song
	err := r.DB.QueryRowContext(context.Background(), queryCheck, id).Scan(&existingSong.ID, &existingSong.Group, &existingSong.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("song not found")
		}
		return nil, err
	}
	
	queryUpdate := "UPDATE songs SET group_name = COALESCE($2, group_name), song_name = COALESCE($3, song_name) WHERE id = $1 RETURNING id, group_name, song_name"
	var updatedSong models.Song
	err = r.DB.QueryRowContext(context.Background(), queryUpdate, id, group, song).Scan(&updatedSong.ID, &updatedSong.Group, &updatedSong.Name)
	if err != nil {
		return nil, err
	}

	return &updatedSong, nil
}

func (r *SongRepository) AddSong(group, song, releaseDate, text, link string) (*models.NewSong, error) {
	query := `
    	INSERT INTO songs (id, group_name, song_name, release_date, lyrics, link)
    	VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)
    	RETURNING id
	`

	var newSong models.NewSong
	err := r.DB.QueryRowContext(context.Background(), query, group, song, releaseDate, text, link).Scan(
		&newSong.ID,
		&newSong.Group,
		&newSong.Song,
		&newSong.ReleaseDate,
		&newSong.Text,
		&newSong.Link,
	)
	if err != nil {
		return nil, err
	}

	return &newSong, nil
}

func SetupDB() *sql.DB {
	conf := getEnv()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME,
	)

	var db *sql.DB
	var err error
	for i := 0; i < retries; i++ {
		db, err = sql.Open("postgres", psqlInfo)

		if err == nil {
			err = db.Ping()

			if err == nil {
				log.Println("Successfully connected to the database.")
				return db
			}
		}

		log.Printf("Retrying to connect to the database (%d/%d): %v", i+1, retries, err)
		time.Sleep(delay)
	}

	log.Fatalf("Failed to connect to the database after %d retries: %v", retries, err)
	return nil
}

func getEnv() models.DBConnection {
	conf := new(models.DBConnection)

	if err := godotenv.Load("/app/config/.env"); err != nil {
		log.Println("Warning: ", err)
	}

	conf.DB_HOST = os.Getenv("DB_HOST")
	conf.DB_PORT = os.Getenv("DB_PORT")
	conf.DB_USER = os.Getenv("DB_USER")
	conf.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	conf.DB_NAME = os.Getenv("DB_NAME")

	if conf.DB_HOST == "" || conf.DB_PORT == "" || conf.DB_USER == "" || conf.DB_PASSWORD == "" || conf.DB_NAME == "" {
		log.Fatal("One or more required database environment variables are missing.")
	}

	return *conf
}
