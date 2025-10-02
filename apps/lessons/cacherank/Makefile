.PHONY: help test test_race coverage coverage_html coverage_clean bench

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test:  ## run the tests
	go test -v ./...

test_race:  ## run the tests with race detector
	go test -v -race ./...

coverage.out:  ## generate coverage profile
	go test -v -coverprofile=coverage.out ./...

coverage: coverage.out  ## run tests with coverage report
	go tool cover -func=coverage.out

coverage_html: coverage.out  ## run tests with coverage and open HTML report
	go tool cover -html=coverage.out

coverage_clean:  ## remove coverage files
	rm -f coverage.out

bench:  ## run benchmark tests
	go test -bench=. -benchmem ./...
