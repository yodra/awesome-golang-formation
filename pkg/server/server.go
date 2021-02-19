package server

import (
	"fmt"
	"github.com/yodra/awesome-golang-formation/pkg/domain"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yodra/awesome-golang-formation/pkg/server/handler/hello"
	"github.com/yodra/awesome-golang-formation/pkg/server/handler/movies"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	movieRepository domain.MovieRepository
}

func New(host string, port uint, movieRepository domain.MovieRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		movieRepository: movieRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("App is up and running on localhost:8080 🎉")
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/hello", hello.Handler())
	s.engine.POST("/movies", movies.CreateHandler(s.movieRepository))
}
