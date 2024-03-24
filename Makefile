.DEFAULT_GOAL := fmt

init:
	rm -rf go.mod
	go mod init github.com/easyone-jwlee/channelizer
	go mod tidy
.PHONY:init

fmt: init
	go fmt ./...
.PHONY:fmt

test: fmt
	go run examples/main.go
.PHONY:test