package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL)
	// fmt.Fprint(w, "Welcome!\n")
}

func withName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	fmt.Println("_____API_SERVICES_STARTED_____")

	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/test/:name", withName)

	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}