.DEFAULT_GOAL := all

LDFLAGS='-extldflags -static'
BIN_NAME=helloworld

all: build

build:
	CGO_ENABLED=0 GOOS=linux go build -o $(BIN_NAME) -a -tags netgo -ldflags $(LDFLAGS) .
.PHONY:build

run:
	go run main.go "/etc/hosts"

clean:
	rm $(BIN_NAME)