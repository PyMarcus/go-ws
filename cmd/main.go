package main

import (
	"log"
	"net/http"
)

func main() {
	r := routes()

	log.Println("Serving at :8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}
