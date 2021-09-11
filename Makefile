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

test-put-read: build
	@mkdir -p tmp
	@echo 'some data' > $(PWD)/tmp/data
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data' -H 'Content-Type: text/html'

test-put-read-single: build
	@mkdir -p tmp
	@echo '{"foo":"bar"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'

test-put-read-multiple: build
	@mkdir -p tmp
	@echo '[{"foo":"bar"}, {"fizz":"buzz"}]' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'

test-put-read-pretty: build
	@mkdir -p tmp
	@echo '{\n  "foo": "bar"\n}' > $(PWD)/tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@tmp/data.json' -H 'Content-Type: application/json'

test-put-read-write: build
	@mkdir -p tmp
	@echo '{"foo":"bar"}' > $(PWD)/tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/put' -X PUT -d '@/tmp/data.json' -H 'Content-Type: application/json' -o tmp/file.txt -s
	@cat $(PWD)/tmp/file.txt

test-post-read-single: build
	@mkdir -p tmp
	@echo '{"data":"foo bar"}' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/post' -X POST -d '@tmp/data.json' -H 'Content-Type: application/json'

test-post-read-multiple: build
	@mkdir -p tmp
	@echo '[{"data":"Thing One"}, {"data":"Thing Two"}]' > $(PWD)/tmp/data.json
	# @tmp/data.json or @./tmp/data.json
	./$(EXEC) -url 'https://httpbin.org/post' -X POST -d '@tmp/data.json' -H 'Content-Type: application/json'

test-head: build
	./$(EXEC) -url 'https://httpbin.org/head' -X HEAD

# Flag testing
test-write-null: build
	./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -s

test-set-headers: build
	@./$(EXEC) -url 'https://httpbin.org/get' -o /dev/null -H "Accept-Language: en,Accept: text/html; application/json,User-Agent: poopoo,Cache-Control: no-cache,Content-Encoding: gzip"
	@make clean

# Package testing
test-httplib: build
	@echo 'go test -v "github.com/bjma/gurl/httplib"

test-filelib: build
	@go test -v "github.com/bjma/gurl/filelib"

test-utils: build
	@echo 'go test -v "github.com/bjma/gurl/utils"
	
