.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY:
docker_build: ## Build docker image
	docker build -t rexposadas/dockerize_example .
	@ echo ✅ success!


.PHONY: run
docker_run: docker_build ## Buld and run run docker images
	docker run -p 8080:8080 rexposadas/dockerize_example
	@ echo ✅ success!