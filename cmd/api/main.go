package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yodra/awesome-golang-formation/pkg/server"
	"github.com/yodra/awesome-golang-formation/pkg/storage/mysql"
	"log"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "leanmind"
	dbPass = "leanmind"
	dbHost = "127.0.0.1"
	dbPort = "3306"
	dbName = "leanmind"
)

func main() {
	mysqlUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlUri)
	if err != nil {
		log.Fatalf("%v %s", err, "cannot connect to the database")
	}
	repo := mysql.NewMovieRepository(db)

	srv := server.New(host, port, repo)
	srv.Run()
}
