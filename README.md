This is a repository of a go stack. This is meant to be a learnig too. 

# Directories and what's in the

## simple_service
A simple REST API. Use this to learn how to: 
1. Create a simple REST API
2. Test the logic of your API

# Running the services locally

## simple_service directory

1. cd to the go directory
2. run `go run main.go`
3. If you want to set the port you can do so with `PORT=3000 go run main.go`
4. call the `/ping` service `curl http://localhost:3000`

You should get a response of 200 with the message `pong`.

You can read more about this directory in the `readme.md` inside the directory.

## Via Docker

    docker build --name rexposadas/simple_service .


# Notes on the different implementations

## simple service

Gin was used because it's one of the more common web frameworks. It's also very easy to use. We could have just as
easily used the standard library, but I wanted to show how to use a framework.


