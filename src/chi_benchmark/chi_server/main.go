package main

import (
	"chi_benchmark"
	"net/http"
)

func main() {

	http.ListenAndServe(":9999", chi_benchmark.NewApi())
}
