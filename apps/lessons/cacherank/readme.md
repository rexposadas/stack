# Simple Cache Implementation

This is a thread-safe cache implementation with the following characteristics:

## Features
- Fixed-size capacity
- Thread-safe operations using mutex synchronization
- Concurrent access support

## Design Constraints
1. The cache has a fixed size defined at initialization
2. `Connection.getItem` always returns a new `Rankable` object
3. Thread-safe operations are guaranteed for all cache methods

## Testing

The implementation includes comprehensive test coverage including:
- Basic cache operations
- Concurrent access testing
- Race condition verification

### Running Tests

To run the tests, run `make test`

To run the tests with the race detector, run `make test_race`. The race detector will check for data races when accessing the cache concurrently.

## Implementation Notes
The cache uses a mutex to ensure thread-safety during concurrent operations. All public methods are protected against race conditions, and the implementation has been verified using Go's race detector.

