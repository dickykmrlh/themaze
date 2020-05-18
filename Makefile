pkgs = $(shell go list ./...)

.PHONY: build

# go build command
build:
	@go build -v -o themaze cmd/*.go

# go run command
run:
	make build
	@./themaze

test:
	@echo "RUN TESTING..."
	@go test -v -cover -race $(pkgs)