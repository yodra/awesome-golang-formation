package server

import "context"

type Movie struct {
	name   string
	year   string
	author string
}

type MovieRepo interface {
	Save(ctx context.Context, movie Movie) error
}

func FormatToDomain(request CreateMovieRequest) Movie {
	return Movie{
		name:   request.Name,
		year:   request.Year,
		author: request.Author,
	}
}