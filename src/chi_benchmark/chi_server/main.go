package main

import (
	"chi_benchmark"
	"net/http"
)

func main() {

	http.ListenAndServe(":8000", chi_benchmark.NewApi())
}
