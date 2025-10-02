# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a learning-focused repository demonstrating how to build applications primarily in Go, with some Python code. The repository is organized to be self-contained, where each directory teaches a specific topic without requiring knowledge of other directories.

## Key Commands

### Running All Services
```bash
make compose      # Start all services via docker-compose
make health       # Verify all services are running
make compose_down # Tear down all services
make test         # Run all Go tests
```

### Working with Individual Services
Each service in `apps/lessons/` has its own Makefile with standard targets:

```bash
# From within a service directory (e.g., apps/lessons/simple_service/)
make test              # Run tests (with optional TEST_ARGS and TEST_TARGETS)
make bench             # Run benchmarks
make docker_build      # Build Docker image (tagged as rexposadas/<service_name>)
make docker_run        # Build and run Docker image
```

**Note**: Before running `make compose` at the root level, you must first run `make docker_build` in each service directory.

### Running Specific Tests
```bash
# From service directory or root
go test -v -run NameOfTest   # Run a specific test
```

### Python Development
The `pystack/` directory contains Python code:
```bash
cd pystack
python -m venv venv          # Create virtual environment
source venv/bin/activate     # Activate environment (macOS/Linux)
pip install -r requirements.txt
```

## Architecture and Structure

### Project Organization

- **`apps/lessons/`**: Self-contained Go learning modules, each teaching a specific concept
  - `simple_service/`: Basic REST API with Gin framework
  - `channels/`: Go channels and concurrency patterns
  - `generics/`: Go generics examples
  - `datastructure/`: Data structure implementations
  - `sync_dir/`: Synchronization primitives
  - `cacherank/`, `lrucache/`: Caching implementations
  - `errgroup/`, `interface/`, `types/`: Various Go concepts

- **`leetcode/`**: LeetCode solutions organized by problem number (e.g., `1_two_sum/`, `121_best_time_to_buy_stock/`)

- **`pystack/`**: Python learning code including:
  - `pydantic_ai_samples/`: Pydantic examples
  - `leetcode/`: Python LeetCode solutions

- **`lib/`**: Shared code reused across projects (minimal dependencies between modules)

- **`apps/cli/`**: CLI tool examples

- **`apps/aoc/`**: Advent of Code solutions

### Learning Patterns

1. **Start Here Files**: Some lessons contain `start_here.go` or `start_here_test.go` files which provide the most basic implementation and introduction to the topic.

2. **Notes Comments**: Throughout the codebase, look for `// Notes:` comments which highlight important learning points and explanations.

3. **Self-Contained Modules**: Each lesson directory is designed to be understood independently without needing to reference other directories.

### Docker Architecture

All Docker images are prefixed with `rexposadas/`. The `docker-compose.yaml` at the root orchestrates services, with each service having its own Dockerfile within its directory.

Services expose health check endpoints (e.g., `http://localhost:8080/health`).

### Go Module Structure

- Module path: `github.com/rexposadas/stack`
- Go version: 1.19
- Key dependencies:
  - `github.com/gin-gonic/gin`: Web framework (implied from service code)
  - `github.com/spf13/cobra`: CLI framework
  - `github.com/stretchr/testify`: Testing utilities
  - `golang.org/x/sync`: Advanced concurrency primitives

### Testing Philosophy

Tests are co-located with source files (e.g., `logic.go` and `logic_test.go`). The codebase includes:
- Unit tests
- Benchmark tests (`make bench`)
- Fuzz tests (where applicable, e.g., `make fuzz` in some services)

## Quick Reference

The `cheat.md` file contains quick reference snippets for common Go patterns including:
- Slice iteration
- Channel operations (buffered/unbuffered, select statements)
- Type assertions and conversions
- Reflection and interfaces
- Benchmarking patterns
