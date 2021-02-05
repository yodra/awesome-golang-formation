package server

type Movie struct {
	name   string
	year   string
	author string
}

type MovieRepo interface {
	Save(movie Movie) error
}
