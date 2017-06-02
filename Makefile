GOPATH=$(shell pwd)
GOBIN=$(GOPATH)/bin
GOPKG=$(GOPATH)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := benchmark

.PHONY: benchmark
benchmark:
	$(GOCMD) test  -bench=".*" -benchtime 1s -benchmem benchmark

.PHONY: dependencies
dependencies:
	$(GOCMD) get -t benchmark

.PHONY: clean
clean:
	rm -f *.log
	rm -fr src/github.com
	rm -fr $(GOBIN)
	rm -fr $(GOPKG)

.PHONY: golax
golax:
	$(GOCMD) get golax_benchmark/golax_server
	$(GOCMD) install golax_benchmark/golax_server
	$(GOBIN)/golax_server >> golax.log 2>&1

.PHONY: gorilla
gorilla:
	$(GOCMD) get gorilla_benchmark/gorilla_server
	$(GOCMD) install gorilla_benchmark/gorilla_server
	$(GOBIN)/gorilla_server >> gorilla.log 2>&1

.PHONY: chi
chi:
	$(GOCMD) get chi_benchmark/chi_server
	$(GOCMD) install chi_benchmark/chi_server
	$(GOBIN)/chi_server >> chi.log 2>&1

