package main

import (
	"log"
	"net/http"
)

func main() {
	route := NewRouter()
	log.Fatal(http.ListenAndServe(":8081", route))
}
