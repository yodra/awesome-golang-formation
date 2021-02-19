package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yodra/awesome-golang-formation/pkg/domain"
	"testing"
)

func Test_Save_RepositoryError(t *testing.T) {
	name, year, author := "Movie name", "2002", "Movie author"
	movie, err := domain.NewMovie(name, year, author)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	sqlMock.ExpectExec(
		"INSERT INTO movies (name, year, author) VALUES (?, ?, ?)").
		WithArgs(name, year, author).
		WillReturnError(errors.New("something-failed"))

	repo := NewMovieRepository(db)

	err = repo.Save(context.Background(), movie)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_Save_Succeed(t *testing.T) {
	name, year, author := "Movie name", "2002", "Movie author"
	movie, err := domain.NewMovie(name, year, author)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	sqlMock.ExpectExec(
		"INSERT INTO movies (name, year, author) VALUES (?, ?, ?)").
		WithArgs(name, year, author).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMovieRepository(db)

	err = repo.Save(context.Background(), movie)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}