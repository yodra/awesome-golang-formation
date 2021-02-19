package server

import (
	"context"
	"errors"
)

type Movie struct {
	name   MovieName
	year   MovieYear
	author MovieAuthor
}

type MovieRepo interface {
	Save(ctx context.Context, movie Movie) error
}

var ErrEmptyField = errors.New("the field can not be empty")

func NewMovie(name, year, author string) (Movie, error) {
	nameVO, err := NewMovieName(name)
	if err != nil {
		return Movie{}, err
	}

	yearVO, err := NewMovieYear(year)
	if err != nil {
		return Movie{}, err
	}

	authorVO, err := NewMovieAuthor(author)
	if err != nil {
		return Movie{}, err
	}

	return Movie{nameVO, yearVO, authorVO}, nil
}

type MovieName struct {
	value string
}

func NewMovieName(name string) (MovieName, error) {
	if name != "" {
		return MovieName{}, ErrEmptyField
	}
	return MovieName{name}, nil
}

type MovieYear struct {
	value string
}

func NewMovieYear(year string) (MovieYear, error) {
	if year != "" {
		return MovieYear{}, ErrEmptyField
	}
	return MovieYear{year}, nil
}

type MovieAuthor struct {
	value string
}

func NewMovieAuthor(author string) (MovieAuthor, error) {
	if author != "" {
		return MovieAuthor{}, ErrEmptyField
	}
	return MovieAuthor{author}, nil
}
