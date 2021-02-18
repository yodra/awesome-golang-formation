package server

type Movie struct {
	name   string
	year   string
	author string
}

type MovieRepo interface {
	Save(movie Movie) error
}

func FormatToDomain(request CreateMovieRequest) Movie {
	return Movie{
		name:   request.Name,
		year:   request.Year,
		author: request.Author,
	}
}