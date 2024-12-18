include .envrc

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: run/api
run/api:
	go run ./cmd/api

.PHONEY: run/dev/api
run/dev/api:
	air

.PHONY: run/install
run/install:
	go build -o ./bin/cli/ cmd/cli/install.go