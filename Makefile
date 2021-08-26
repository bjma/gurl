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
	@rm -rf $(PWD)/tmp

# shitty way of testing rn bc i dont feel like reading Go's testing library
test-get: build
	./$(EXEC) 'https://httpbin.org/get'

test-put: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json'

test-get-write: build
	./$(EXEC) -url 'https://httpbin.org/get' -o @file.txt -s
	@cat $(PWD)/tmp/file.txt