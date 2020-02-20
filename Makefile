#GOBIN = $(shell pwd)/build/bin
GOBIN = $(shell pwd)
GO ?= latest

.PHONY : vote

vote:
	go build -o vote ./server/
	@echo "Done building."
	@echo "Run \"$(GOBIN)/vote\" to launch vote server."