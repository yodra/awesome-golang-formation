package main

import (
	"github.com/gorilla/mux"
	"github.com/yodra/awesome-golang-formation/server/handler/hello"
	"log"
	"net/http"
)

func main() {

	routerParent := mux.NewRouter().StrictSlash(true)

	routerParent.HandleFunc("/hello", hello.Handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", routerParent))
}
