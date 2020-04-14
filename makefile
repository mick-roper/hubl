.PHONY: all

all: clean install test build

clean:
	rm -rf bin/

install:
	go get ./src/...

test:
	go test ./src/...

build:
	GOARCH=amd64 GOOS=linux go build -o bin/linux/hubl src/main.go