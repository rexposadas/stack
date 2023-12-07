.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: test
TEST_ARGS ?= -v
TEST_TARGETS ?= ./...
test: ## Test the Go modules within this package.
	@ echo ▶️ go test $(TEST_ARGS) $(TEST_TARGETS)
	go test $(TEST_ARGS) $(TEST_TARGETS)
	@ echo ✅ success!


.PHONY: bench
BENCH_ARGS ?= -bench
BENCH_TARGETS ?=.
bench: ## Run benchmarks
	@ echo ▶️ go test $(BENCH_ARGS) $(BENCH_TARGETS)
	go test $(BENCH_ARGS) $(BENCH_TARGETS)
	@ echo ✅ success!


.PHONY: fuzz
FUZZ_ARGS ?= -fuzz
FUZZ_TARGETS ?=FuzzAdd -fuzztime 5s
fuzz: ## Run Fuzz on func Add()
	@ echo ▶️ go test $(FUZZ_ARGS) $(FUZZ_TARGETS) 10s
	go test $(FUZZ_ARGS) $(FUZZ_TARGETS)
	@ echo ✅ success!

.PHONY: docker_build
docker_build: ## Build docker image
	docker build -t rexposadas/simple_service .


.PHONY: docker_run
docker_run: docker_build ## Build and run run docker image
	docker run -p 8080:8080 rexposadas/simple_service