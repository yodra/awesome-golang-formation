package mysql

import (
	"context"
	"database/sql"
	"github.com/yodra/awesome-golang-formation/server"
)

type MovieRepository struct {
	db *sql.DB
}

func (repo MovieRepository) Save(ctx context.Context, movie server.Movie) error {
	return nil
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}
