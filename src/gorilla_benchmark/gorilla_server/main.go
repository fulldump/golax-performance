package main

import (
	"gorilla_benchmark"
	"log"
	"net/http"
)

func main() {

	router := gorilla_benchmark.NewApi()

	log.Fatal(http.ListenAndServe(":9999", router))
}
