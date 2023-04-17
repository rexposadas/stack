This is a repository of a go stack. 

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

## dockerize
Teaches how to:
1. Create a docker image from a service.
2. Run a docker container.
3. Make an HTTP call to a running container.


# Running the services

Instructions on how to work with the individual services are in the `readme.md` inside the directory.


