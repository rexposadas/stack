test:
	go test ./...

bench:
	go test -bench=.

fuzz:
	# limit fuzz test to 5 seconds.
	# To run it indefinitely, remove the -fuzztime flag.
	go test -fuzz=FuzzAdd -fuzztime 5s

all: test bench