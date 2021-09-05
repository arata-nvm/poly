.PHONY: build
build:
	go build -o poly ./examples

.PHONY: run
run:
	go run ./examples
