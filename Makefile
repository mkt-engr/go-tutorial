.DEFAULT_GOAL := build

.PHONY: fmt vet build clean run

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	mkdir -p bin
	go build -o bin/

run: build
	./bin/*

clean:
	rm -rf bin
	go clean