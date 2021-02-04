package main

import (
	"awesome-golang-formation/server/handler/hello"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	routerParent := mux.NewRouter().StrictSlash(true)

	routerParent.HandleFunc("/hello", hello.Handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", routerParent))
}
