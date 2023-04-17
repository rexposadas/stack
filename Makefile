.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)



.PHONY: compose_down
compose_down: ## Tear down services
	docker compose down


.PHONY: compose
compose: compose_down ## Test the Go modules within this package.
	docker compose up --remove-orphans


.PHONY: health
health: ## Make a call to the health services to ensure that the services are up and running.
	@echo calling simple service
	curl http://localhost:8080/health

	@echo calling dockerize service
	curl http://localhost:8081/health
