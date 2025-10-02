# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a learning module demonstrating thread-safe cache implementation in Go. The cache uses a rank-based eviction policy, where when capacity is reached, the lowest-ranking item is removed before adding new items.

## Key Commands

### Testing
```bash
make test       # Run all tests
make test_race  # Run tests with race detector (CRITICAL for concurrent code)
```

### Running Specific Tests
```bash
go test -v -run TestCacheConcurrentAccess  # Test concurrent access
go test -v -run TestCacheConcurrentCapacity  # Test capacity under concurrency
```

## Architecture and Key Concepts

### Core Components

**Cache (`cache.go`)**: Thread-safe cache with fixed capacity
- Uses `sync.Mutex` to protect all operations
- Maps string keys to `Rankable` items
- Evicts lowest-ranking item when at capacity
- All public methods (`Get`, `Size`) acquire mutex lock

**Rankable (`cache.go`)**: Simple struct with `key` (string) and `rank` (int)

**Connection (`connection.go`)**: Mock data source that returns `Rankable` items with random ranks (1-100)

### Critical Design Constraints

1. **Fixed Capacity**: Cache size is set at initialization and never changes
2. **Rank-Based Eviction**: When full, the item with the lowest rank is removed
3. **Thread Safety**: All cache operations are protected by mutex
4. **Connection Behavior**: `Connection.getItem()` always returns a new `Rankable` object with random rank

### Concurrency Patterns

The implementation demonstrates proper Go concurrency practices:

- **Mutex Protection**: All shared state access is wrapped in `c.mutex.Lock()/Unlock()`
- **Defer Unlock**: Uses `defer c.mutex.Unlock()` pattern to ensure mutex is released even if panic occurs
- **Race Detection**: Tests include concurrent access scenarios; always run with `-race` flag when modifying

### Key Implementation Details

**Eviction Algorithm** (`removesLowestRankingItem`):
1. Iterates through all cached items
2. Finds item with minimum rank value
3. Deletes that item from the map

**Cache Get Flow** (`Get`):
1. Lock mutex
2. Check if item exists in cache → return if found
3. Fetch from Connection (always creates new item)
4. If under capacity → add item and return
5. If at capacity → remove lowest rank, add new item, return

### Testing Strategy

Tests cover three critical scenarios:
1. **Basic Operations**: Single-threaded add/get/eviction
2. **Concurrent Access**: Multiple goroutines accessing same key simultaneously
3. **Concurrent Capacity**: Verify capacity constraint holds under concurrent writes

**Important**: Always run `make test_race` when modifying cache logic. Data races are subtle and may not appear in standard tests.

## Module Context

This is part of the larger `github.com/rexposadas/stack` learning repository. The cacherank module is self-contained and can be understood without knowledge of other modules in the stack.
