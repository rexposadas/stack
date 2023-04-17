# Purpose

This directory shows how to dockerize a service and work with a running container.å

# Notes

Since this project is meant to teach a topic in Go, I peppered it with comments starting with `// Notes`.

These notes indicates what I wanted to teach within this particular project.


# Running the app

In a terminal, build the image rexposadas/dockerize_example: 

    make docker_build 


The command above creates an image in the docker registery in your localhost.  If you run 

    docker images

you will see the following image: `rexposadas/dockerize_example`


Run the container

    make docker_run


In another terminal, make an HTTP call to the running container:
    
    curl http://localhost:8000/


You should get an HTTP 200 reponse