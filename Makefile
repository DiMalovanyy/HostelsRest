.PHONY: build
build:
	go build -v ./cmd/restapi

.PHONY: clean
clean:
	rm ./restapi

.DEFAULT_GOAL := build