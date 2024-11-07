.DEFAULT_TARGET: build

.PHONY: fmt vet build

run: build
	@./bin/app

build: vet
	@templ generate
	@go build -o ./bin/app .

vet: fmt
	@go vet ./...

fmt:
	@go fmt ./...