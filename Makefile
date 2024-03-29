include .envrc

# ================================ #
# HELP
# ================================ #
.PHONY: help 
help:
	go run ./cmd/api --help

# ================================ #
# DEVELOPMENT
# ================================ #
.PHONY: run
run:
	go run ./cmd/api --timeout=${timeout} --filename=${filename} --shuffle=${shuffle}

# ================================ #
# QUALITY CONTROL
# ================================ #
.PHONY: audit
audit: 
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Linting code...'
	staticcheck ./...

# ================================ #
# Build
# ================================ #
.PHONY: build
buld:
	go build -o=./bin/api ./cmd/api


