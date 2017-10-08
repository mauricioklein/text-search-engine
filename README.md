[![Build Status](https://travis-ci.org/mauricioklein/text-search-engine.svg?branch=master)](https://travis-ci.org/mauricioklein/text-search-engine)
[![Coverage Status](https://coveralls.io/repos/github/mauricioklein/text-search-engine/badge.svg?branch=master)](https://coveralls.io/github/mauricioklein/text-search-engine?branch=master)

# text-search-engine
A Golang driven text search engine

## Dependencies

- Go 1.9
- Govendor 1.0.8 or superior

## Setup

```bash
# Download external dependencies
$ govendor sync

# Build the project and install the binary
$ govendor install +local
````

## Specs

To run the specs:

```bash
$ govendor test -v +local
```

## Run

To run the project:

```bash
# Run the binary
$ $GOPATH/bin/text-search-engine -directory [file's directory]
```

To check all the available command line parameters, run:

```bash
$ $GOPATH/bin/text-search-engine -h
```

## Terminal

Right after running the system, an interactive terminal will be present on the console.
This console will accept any character as input until the first line break.
For each input sentence, the rank is calculated against each read file and the top 10 ranks will be displayed
as result.
To quit the system, just submit the quit sentence: `:quit`

```bash
$ ./text-search-engine -directory test-utils/11-files

11 file(s) read in the directory test-utils/11-files

search> Lorem
file1.txt: 100.00% match
file10.txt: 100.00% match
file11.txt: 100.00% match
file3.txt: 100.00% match
file6.txt: 100.00% match
file7.txt: 100.00% match
file8.txt: 100.00% match
file9.txt: 100.00% match
file2.txt: 0.00% match
file4.txt: 0.00% match

search> dolor
file1.txt: 100.00% match
file10.txt: 100.00% match
file11.txt: 100.00% match
file2.txt: 100.00% match
file3.txt: 100.00% match
file4.txt: 100.00% match
file6.txt: 100.00% match
file7.txt: 100.00% match
file8.txt: 100.00% match
file9.txt: 100.00% match

search> :quit 

$
```

## The Algorithm

The matching algorithm implemented is based on the [Levenshtein distance algorithm](https://en.wikipedia.org/wiki/Levenshtein_distance).

The input sequence is broken into words. A word is defined by any sequence of characters, delimited
by a space or a line break. Thus, for each word, the Levenshtein distance to the text is calculated.
The score is based on the shortest distance calculated.
Finally, the final score is the average value among all the calculated scores

Example:

- Input sentence: "Far Fo"
- Content: "The Far Fox"

Distances:

| Sentence Word | File Word | Distance    |
|:-------------:|:---------:|:-----------:|
| Far           | The       | 3           |
| Far           | Far       | 0 (minimum) |
| Far           | Fox       | 2           |
| Fo            | The       | 3           |
| Fo            | Far       | 2           |
| Fo            | Fox       | 1 (minimum) |

-  Score for "Far": 0 (minimum distance) / 3 letters of word ->  *Score: 1.000*
-  Score for "Fo": 1 (minimum distance) / 2 letters of word ->  *Score: 0.666*


> Rank = (1.0 + 0.666) / 2 = 0.833 or **83.3% Match**

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
