[![Build Status](https://travis-ci.org/mauricioklein/text-search-engine.svg?branch=master)](https://travis-ci.org/mauricioklein/text-search-engine)
[![Coverage Status](https://coveralls.io/repos/github/mauricioklein/text-search-engine/badge.svg?branch=master)](https://coveralls.io/github/mauricioklein/text-search-engine?branch=master)

# text-search-engine
A Golang driven text search engine

## Docker

This project provides a `Dockerfile`, which allows to create a Docker image with the application installed
and ready to use.

### Setup

To create the docker image, run:

```bash
$ docker build -t text-search-engine .
```

A new Docker image called `text-search-engine` will be created at the end of the process.

### Run

To run the system in a container, first we need to map to the container the directory with the files that will
be processed.

As an example, let's consider we've a directory with files saved in `/home/foobar/files`.
The Docker command will be:

```bash
$ docker run -v /home/foobar/files:/tmp/files -it text-search-engine -directory /tmp/files
```

The parameter `-v` indicates to map the localhost directory `/home/foobar/files` to the container directory
`/tmp/files`.

The parameter `-it` makes the container run in interactive mode, allowing to interact with container's STDIN.

Finally, the last parameter `-directory /tmp/files` indicate the file's directory to the search engine app.
