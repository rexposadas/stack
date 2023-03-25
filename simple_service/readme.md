# Purpose of this directory

This is a demonstration of how to create a REST service using 
gin and simple basic stuff like testing and benchmarcking. 

You should be able to look at the code and see how things are done easily
since this are the simple cases. Start by looking at `main.go`


# Getting started

## Running tests
Run `make test` to run the tests

Or simply `go test ./..`

## Calling the service

In one terminal: `PORT=3000 go run main.go`

In another terminal: `curl http://localhost:3000/ping`

# What's here

1. A simple REST service using gin
2. Look at logic.go for :
   3. Example of testing
   4. Examples of benchmarking
5. Makefile gives a simple way to run the test