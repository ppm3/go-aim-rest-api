# Makefile
MODULE = go-aim-rest-api
GOCMD=go
BINARY_UNIX=$(MODULE)_unix
GOGET=$(GOCMD) get
PROJECT = $(GOPROJECTS)$(MODULE)
GOCOVER=$(GOCMD) tool cover

.PHONY: generate run build test lint deps vendor test/cover

generate:
	$(GOCMD)  generate ./...

clean:
	rm -rf ./bin
	rm -rf ./out

run:
	export ENVIRONMENT=development && \
	$(GOCMD)  run $(PROJECT)/cmd/server/main.go

build: # build a server
	$(GOCMD)  build -a -o $(MODULE) $(PROJECT)/build

test:
	$(GOCMD)  clean -testcache
	$(GOCMD)  test ./... -v

lint:
	gofmt -l .

vendor:
	$(GOCMD)  mod vendor

test/coverage:
	$(GOCMD)  clean -testcache
	$(GOCMD)  test -v -coverprofile=coverage.out ./...
	$(GOCOVER)  -func=coverage.out
	$(GOCOVER)  -html=coverage.out

deps:
	$(GOGET)  github.com/gin-gonic/gin
	$(GOGET)  go.mongodb.org/mongo-driver/mongo
	$(GOGET)  go.mongodb.org/mongo-driver/mongo/options
	$(GOGET)  github.com/go-sql-driver/mysql
	$(GOGET)  github.com/joho/godotenv
	$(GOGET)  github.com/stretchr/testify/assert
	$(GOGET)  github.com/stretchr/testify/mock