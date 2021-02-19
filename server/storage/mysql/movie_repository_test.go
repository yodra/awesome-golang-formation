package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yodra/awesome-golang-formation/server"
	"testing"
)

func Test_Save_RepositoryError(t *testing.T) {
	name, year, author := "Movie name", "2002", "Movie author"
	movie, err := server.NewMovie(name, year, author)
	if err != nil {
		t.Fatal(err)
	}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	sqlMock.ExpectExec(
		"INSERT INTO movies (name, year, author) VALUES (?, ?, ?)").
		WithArgs(name, year, author).
		WillReturnError(errors.New("something-failed"))

	repo := NewMovieRepository(db)

	err = repo.Save(context.Background(), movie)

}

func Test_Save_Succeed(t *testing.T) {
	name, year, author := "Movie name", "2002", "Movie author"
	movie, err := server.NewMovie(name, year, author)
	if err != nil {
		t.Fatal(err)
	}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(err)
	}
	sqlMock.ExpectExec(
		"INSERT INTO movies (name, year, author) VALUES (?, ?, ?)").
		WithArgs(name, year, author).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMovieRepository(db)

	err = repo.Save(context.Background(), movie)

}