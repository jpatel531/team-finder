package main

import (
	"github.com/jpatel531/team-finder/search"
	"log"
	"sort"
)

const (
	endpoint          = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
	batchSize         = 100
	startID           = 1
	increment         = 1
	notFoundThreshold = 100
)

func main() {
	targets := []string{
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester Utd",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"FC Bayern Munich",
	}

	s := search.New(endpoint, targets)

	complete := s.DoBatch(startID, batchSize, increment, notFoundThreshold)
	if complete {
		log.Println("All teams found 🎉")

		players := s.PlayersFound()
		sort.Sort(players)
		log.Println(players)
	} else {
		log.Println("Could not find teams 😱")
	}
}
