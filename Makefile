.PHONY: watch
watch: ## Start a file watcher to run tests on change. (requires: watchexec)
	watchexec -c "go test -timeout 1m -failfast ./..."

.PHONY: all
all: lint test

.PHONY: test
test: ## Runs the unit test suite
	go test -timeout 1m -race ./...

.PHONY: lint
lint: ## Runs the linters (including internal ones)
	golint -set_exit_status ./...;
	errcheck ./...;
	gosec -quiet ./...;
	staticcheck ./...;

.PHONY: fuzz
fuzz: ## Runs fuzzing (won't terminate on its own)
	go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
	go-fuzz-build
	go-fuzz -workdir testdata/fuzz

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.

.DEFAULT_GOAL := help

help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'