# Team Finder

A program that uses a REST API to find certain football teams by name. For every player in those football teams, sort them alphabetically by name and print their full name, age and list of teams.

## Installation

To place this package in your `$GOPATH`, you can use `go get`:

	$ go get github.com/jpatel531/team-finder

## Execute The Program

Go into the project directory, build the package and run the output binary.

	$ cd $GOPATH/src/github.com/jpatel531/team-finder
	$ go build && ./team-finder

To see how long the program runs for:

	$ /usr/bin/time -l ./team-finder

## Running The Tests

Ensure that you have these dependencies installed, both from [testify/assert](https://github.com/stretchr/testify):

* github.com/stretchr/testify/assert
* github.com/stretchr/testify/mock

In the root of the project, you can run:

	$ go test -v ./...