package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/yodra/awesome-golang-formation/server"
)

type MovieRepository struct {
	db *sql.DB
}

func (repo MovieRepository) Save(ctx context.Context, movie server.Movie) error {
	insert, err := repo.db.Prepare("INSERT INTO movies (name, year,author) VALUES (?,?,?)")
	if err != nil {
		return err
	}

	_, err = insert.ExecContext(ctx, movie.Name(), movie.Year(), movie.Author())
	if err != nil {
		return fmt.Errorf("error trying to persist movies on database: %v", err)
	}

	return nil
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}
