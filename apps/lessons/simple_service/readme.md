# Purpose of this directory

This is a demonstration of how to create a REST service using 
gin and simple basic stuff like testing and benchmarcking. 

You should be able to look at the code and see how things are done easily
since this are the simple cases. Start by looking at `main.go`

**What's here**

1. A simple REST service using the gin framework.
2. Look at logic.go for:
   1. Example of testing, benchmarking and fuzzing.
   2. Defining a function that your router will use.
3. Makefile gives a simple way to run the test

# Getting started

The simplest way to get started is to run the `make` targets.

`make help` shows you possible targets.

`make test` to run the tests.

`make bench` to run benchmarks.

The code starts in main.go where the REST API is initialized.

## Calling the service

In one terminal: `PORT=3000 go run main.go`

In another terminal: `curl http://localhost:3000/ping`

## Using Docker

The image was taken from [here](https://hub.docker.com/_/golang)

In a terminal: 

   docker build -t rexposadas/simple_service .
   docker run -p 8080:8080 -it --rm --name simple_service rexposadas/simple_service

The `-p` flag exposes 8080 to the host so that can call:
   
   curl http://localhost:8080/ping




that will build the docker image



