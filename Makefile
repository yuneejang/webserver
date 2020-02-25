GOBIN = $(shell pwd)/build/bin
#GOBIN = $(shell pwd)
GO ?= latest

.PHONY : server client

all : server client

server:
#	go build -o server ./server
	go1.13.8 build -o server ./server
	mv ./server/server $(GOBIN)/
	@echo "Done building."
	@echo "Run \"$(GOBIN)/server\" to launch server."

client:
#	go build -o client ./client
	go1.13.8 build -o client ./client
	mv ./client/client $(GOBIN)/
	@echo "Done building."
	@echo "Run \"$(GOBIN)/client\" to launch client."


