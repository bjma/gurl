-include .env

EXEC=gurl
PWD=$(shell pwd)

# Go variables
GOBIN := $(PWD)/bin

all: build

build:
	go build -o $(EXEC) main.go

run:
	go run main.go

clean:
	@rm -rf $(EXEC)
	@rm -rf $(PWD)/tmp

fmt:
	@go fmt ./...

# shitty way of testing rn bc i dont feel like reading Go's testing library
get: build
	./$(EXEC) 'https://httpbin.org/get'

get-write: build
	./$(EXEC) -url 'https://httpbin.org/get' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt

put: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json'

put-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt

put-read-single: build
	@mkdir -p tmp
	@echo '{"bogos":"binted"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'


put-read-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@/tmp/data.json' -H 'Content-Type: application/json' -o file.txt -s

write-null: build
	./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null

# smaller unit tests
test-util: build
	@go test -v "github.com/bjma/gurl/util"
	