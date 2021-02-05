package mysql

import (
	"database/sql"
	"github.com/yodra/awesome-golang-formation/server"
)

type MovieRepository struct {
	db *sql.DB
}

func (repo MovieRepository) Save(movie server.Movie) error {
	return nil
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}
