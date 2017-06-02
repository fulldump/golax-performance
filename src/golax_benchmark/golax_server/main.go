package main

import (
	"golax_benchmark"
	"net/http"
)

func main() {

	s := &http.Server{
		Addr:    "0.0.0.0:9999",
		Handler: golax_benchmark.NewApi(),
	}

	s.ListenAndServe()

}
