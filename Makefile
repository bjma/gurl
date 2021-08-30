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

# HTTP testing
test-get: build
	./$(EXEC) -url 'https://httpbin.org/get'

test-get-write: build
	./$(EXEC) -url 'https://httpbin.org/get' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt

test-put: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json'

test-put-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt

test-put-read-single: build
	@mkdir -p tmp
	@echo '{"bogos":"binted"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'


test-put-read-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@/tmp/data.json' -H 'Content-Type: application/json' -o file.txt -s

# Flag testing
test-write-null: build
	./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -s

test-set-headers: build
	@./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -H "Accept-Language: en,Accept: text/html; application/json,User-Agent: poopoo,Cache-Control: no-cache,Content-Encoding: gzip"

# Package testing
test-util: build
	@go test -v "github.com/bjma/gurl/util"

test-httplib: build
	@echo 'go test -v "github.com/bjma/gurl/httplib"

test-filelib: build
	@echo 'go test -v "github.com/bjma/gurl/filelib"
	
