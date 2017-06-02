GOPATH=$(shell pwd)
GOBIN=$(GOPATH)/bin
GOPKG=$(GOPATH)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := benchmark

.PHONY: benchmark clean golax gorilla chi

benchmark:
	$(GOCMD) test  -bench=".*" -benchtime 10s -benchmem benchmark

clean:
	rm -f *.log
	rm -fr src/github.com
	rm -fr $(GOBIN)
	rm -fr $(GOPKG)

golax:
	$(GOCMD) get golax_benchmark/golax_server
	$(GOCMD) install golax_benchmark/golax_server
	$(GOBIN)/golax_server >> golax.log 2>&1

gorilla:
	$(GOCMD) get gorilla_benchmark/gorilla_server
	$(GOCMD) install gorilla_benchmark/gorilla_server
	$(GOBIN)/gorilla_server >> gorilla.log 2>&1

chi:
	$(GOCMD) get chi_benchmark/chi_server
	$(GOCMD) install chi_benchmark/chi_server
	$(GOBIN)/chi_server >> chi.log 2>&1

