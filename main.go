package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("GO REST API INSIDE MAIN")
	route := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", route))
	log.Printf("GO REST API STARTED")
}
