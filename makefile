.PHONY: all

all: clean install test build

clean:
	rm -rf bin/

install:
	go get ./src/...

test:
	go test ./src/...

build: build-linux build-windows

build-linux:
	GOARCH=amd64 GOOS=linux go build -o bin/linux/hubl main.go

build-windows:
	GOARCH=amd64 GOOS=windows go build -o bin/windows/hubl.exe main.go