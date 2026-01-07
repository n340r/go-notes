.PHONY: help run run-cron test fmt

help:
	@printf "%s\n" \
		"Targets:" \
		"  make run                         Run with sample cron" \
		"  make run-cron CRON='...expr...'  Run with provided cron" \
		"  make test                        Run tests" \
		"  make fmt                         gofmt ./cmd"

run:
	@go run ./cmd/main.go '*/11 0 1,15 * 1-5 /usr/bin/find'

run-cron:
	@if [ -z "$(CRON)" ]; then \
		echo "Usage: make run-cron CRON='*/15 0 1,15 * 1-5 /usr/bin/find'"; \
		exit 2; \
	fi
	@go run ./cmd/main.go "$(CRON)"

test:
	@go test ./...

fmt:
	@gofmt -w ./cmd/*.go
