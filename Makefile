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
	@make clean

test-get-write: build
	./$(EXEC) -url 'https://httpbin.org/get' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt
	@make clean

test-put: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json'
	@make clean

test-put-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '{"bogos":"binted"}' -H 'Content-Type: application/json' -o tmp/file.txt -s -v
	@cat $(PWD)/tmp/file.txt
	@make clean

test-put-read: build
	@mkdir -p tmp
	@echo 'some data' > $(PWD)/tmp/data
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data' -H 'Content-Type: text/html'
	@make clean

test-put-read-single: build
	@mkdir -p tmp
	@echo '{"data":"you got pwned!"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'http://localhost:8080/api/test' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'
	@make clean

test-put-read-pretty: build
	@mkdir -p tmp
	@echo '{\n  "foo": "bar"\n}' > $(PWD)/tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'
	@make clean

test-put-read-write: build
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@/tmp/data.json' -H 'Content-Type: application/json' -o file.txt -s
	@make clean

test-post-read: build
	@mkdir -p tmp
	@echo '{"data":"you got pwned!"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'http://localhost:8080/api/test' -X POST -d '@tmp/data.json' -H 'Content-Type: application/json'
	@make clean

# Flag testing
test-write-null: build
	./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -s
	@make clean

test-set-headers: build
	@./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -H "Accept-Language: en,Accept: text/html; application/json,User-Agent: poopoo,Cache-Control: no-cache,Content-Encoding: gzip"
	@make clean

# Package testing
test-httplib: build
	@echo 'go test -v "github.com/bjma/gurl/httplib"
	@make clean

test-filelib: build
	@go test -v "github.com/bjma/gurl/filelib"
	@make clean

test-utils: build
	@echo 'go test -v "github.com/bjma/gurl/utils"
	@make clean
	
