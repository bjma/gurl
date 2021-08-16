-include .env

EXEC=gurl

# Go variables
GOBIN := $(shell pwd)/bin

all: build

build:
	go build -o $(EXEC) main.go

run:
	go run main.go

clean:
	@rm -rf $(EXEC)