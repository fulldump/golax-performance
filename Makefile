GOPATH=$(shell pwd)
GOBIN=$(GOPATH)/bin
GOPKG=$(GOPATH)/pkg
GO=go
GOCMD=GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO)

.DEFAULT_GOAL := test

.PHONY: clean golax gorilla

clean:
	rm *.log
	rm -fr src/github.com
	rm -fr $(GOBIN)
	rm -ft $(GOPKG)

golax:
	$(GOCMD) get golax
	$(GOCMD) install golax
	$(GOBIN)/golax >> golax.log 2>&1

gorilla:
	$(GOCMD) get gorilla
	$(GOCMD) install gorilla
	$(GOBIN)/gorilla >> gorilla.log 2>&1
