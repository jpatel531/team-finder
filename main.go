package main

import (
	"github.com/jpatel531/team-finder/search"
	"log"
	"sort"
)

const endpoint = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"

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

	search := search.New(endpoint, targets)

	complete := search.DoBatch(1, 100, 1)
	if complete {
		log.Println("All teams found 🎉")
		players := search.PlayersFound()
		sort.Sort(players)
		log.Println(players)
	} else {
		log.Println("Could not find teams 😱")
	}
}
