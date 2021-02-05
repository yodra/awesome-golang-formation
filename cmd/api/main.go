package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yodra/awesome-golang-formation/server/handler/hello"
	"github.com/yodra/awesome-golang-formation/server/handler/movies"
	"github.com/yodra/awesome-golang-formation/server/storage/mysql"
	"log"
	"net/http"
)

const (
	dbUser = "leanmind"
	dbPass = "leanmind"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "leanmind"
)

func main() {

	routerParent := mux.NewRouter().StrictSlash(true)

	mysqlUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlUri)
	if err != nil {
		log.Fatalf("%v %s", err, "cannot connect to the database")
	}
	repo := mysql.NewMovieRepository(db)

	routerParent.HandleFunc("/hello", hello.Handler).Methods(http.MethodGet)
	routerParent.HandleFunc("/movies", movies.CreateHandler(repo)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", routerParent))
}
