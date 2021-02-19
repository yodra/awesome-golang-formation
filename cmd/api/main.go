package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/yodra/awesome-golang-formation/server/handler/hello"
	"github.com/yodra/awesome-golang-formation/server/handler/movies"
	"github.com/yodra/awesome-golang-formation/server/storage/mysql"
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

	fmt.Println("App is up and running on localhost:8080 🎉")
	log.Fatal(http.ListenAndServe(":8080", routerParent))
}
