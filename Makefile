.DEFAULT_GOAL := fmt

fmt:
	go fmt ./...
.PHONY:fmt

test: fmt
	go run examples/main.go
.PHONY:test