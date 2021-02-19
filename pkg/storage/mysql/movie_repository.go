package mysql

import (
	"context"
	"database/sql"
	"github.com/yodra/awesome-golang-formation/pkg/domain"
)

type MovieRepository struct {
	db *sql.DB
}

func (repo MovieRepository) Save(ctx context.Context, movie domain.Movie) error {
	_, err := repo.db.Exec("INSERT INTO movies (name, year, author) VALUES (?, ?, ?)",
		movie.Name(), movie.Year(), movie.Author())
	if err != nil {
		return err
	}

	return nil
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}
