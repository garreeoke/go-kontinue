package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("go-kontinue is Ready to receive requests")
	log.Fatal(http.ListenAndServe(":8080", router))
}
