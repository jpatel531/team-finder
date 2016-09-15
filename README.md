# Team Finder

A program that uses a REST API to find certain football teams by name. For every player in those football teams, sort them alphabetically by name and print their full name, age and list of teams.

## Challenges

* In the REST API there are only two ways of traversing the list:
	* Increment the `id` route param in the path. `O(n)`.
	* Check a team's previously-played, and next two to-play opponents and follow those paths in essentially a graph. Also `O(n)`, though potentially quicker in some cases (e.g. when you want to search only within a league or upcoming competition).
* Teams are ordered only by ID and not by name.
* Team count is unknown.

## Decisions

Although I was very interested in the idea of traversing the graph of teams and recent opponents, I realized this could cause a lot of wasted requests as the clubs we may be looking for could be in completely disparate leagues. The graph approach tends to favour searching within a team's league or competitions. It could lead to dead-ends, and also - given that one would say in a club football graph or an international football graph - it alone couldn't be enough.

I decided to go for more list-like approach of incrementing IDs and trying each page, but on each search-miss, checking the previous, next, and following opponents to see if we can make a 'jump' to a targeted team.

With regards to speed and performance, I opted to search the API in ranges, such as IDs 1-100. These 100 requests would be executed concurrently and the program would block until this search batch is done. If the search is complete (all teams have been found), it prints out the players. If not, it tackles batches 101-200. To stop this occuring forever if the team does not exist in the API, if the number of 404s reaches a threshold, the program exits.

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