This is a repository demonstrates how to build "things" using Go.  Each directory is meant to be a learning tool to help you learn how to build things with go.  Currently, this has content to teach the following:

1 - Creating web services and running them using `docker`.

2 - Working `docker compose` 

3 - Working with generics.

# Goals of the project

1 - You should be able to pick a top you want to learn, go the directory where that's taught and run + learn everything inside that directory. Each directory should be self-contained. There should be no need to look at files from a different directory. For example. If I wanted to learn how to `dockerize` a webservice, I go to the `dockerize` directory and learn how to do it from there.

2 - Demonstrate how to work with different go technologies.

# Learning

In each of the directories I have comments which starts with `// Notes`. These indicate 
the things I want to point out in each project.


# Getting started

### tldr:

In one terminal, start the services:

    make compose 

In another terminal, ensure that the services are up:

    make health


### Breaking things down further on how to get started 

**Caveat**: You need to go to each service directory and do `make docker_build` before you can do the steps below. It's a bug that will get fixed.

You can get a feel for what this project does by starting all the services.

    make compose

runs the docker compose and starts the services. You can see the docker containers that gets started.  All docker containers are prefixed with `rexposadas`
    
    docker ps 

You can then check the health of the services.

    make health

# Directories and what's in them

## simple_service
A simple REST API. Use this to learn how to: 
1. Create a simple REST API
2. Test the logic of your API

# Running the services

Instructions on how to work with the individual services are in the `readme.md` inside the directory.


