# HelloService

This is meant to be a learning tool.

# Directories and what's in them

## go
This is a go implementation 


# Running the services locally

## simple_service directory

1. cd to the go directory
2. run `go run main.go`
3. If you want to set the port you can do so with `PORT=3000 go run main.go`
4. call the `/ping` service `curl http://localhost:3000`

You should get a response of 200 with the message `pong`.

You can read more about this directory in the `readme.md` inside the directory.

# Notes on the different implementations

## Go

Gin was used because it's one of the more common web frameworks. It's also very easy to use. We could have just as 
easily used the standard library, but I wanted to show how to use a framework.

