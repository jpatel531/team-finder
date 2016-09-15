package targets

import (
	"github.com/jpatel531/team-finder/teams"
	"log"
	"sync"
)

type Targets interface {
	TeamIsTargeted(teamName string) bool
	Hit(team *teams.Team)
	IsComplete() bool
	PlayersFound() teams.Players
}

func New(teamNames []string) Targets {
	targetedTeams := make(map[string]*teams.Team)
	for _, name := range teamNames {
		targetedTeams[name] = nil
	}
	return &targets{
		teams:   targetedTeams,
		RWMutex: &sync.RWMutex{},
		players: make(map[string]*teams.Player),
	}
}

type targets struct {
	*sync.RWMutex
	teams   map[string]*teams.Team
	players map[string]*teams.Player
}

func (t *targets) TeamIsTargeted(teamName string) bool {
	t.RLock()
	defer t.RUnlock()
	team, exists := t.teams[teamName]
	return exists && team == nil
}

func (t *targets) Hit(team *teams.Team) {
	name := team.Name
	if !t.TeamIsTargeted(name) {
		log.Panicf("%s is not a target in %+v\n\n", name, t.teams)
	}
	t.Lock()
	defer t.Unlock()
	t.teams[name] = team
	for _, player := range team.Players {
		// the club listing of the player contains the more complete information
		// about a player's teams. That is, it shows both club and country.
		// Therefore, we should prioritize the club listing over the national listing.
		if !team.IsNational {
			player.Club = team.Name
		}

		// Always write the player to the map if the team is a club team.
		// Otherwise, write it if it's nil. If the player also plays for one of the target clubs
		// the map will be overwritten with the club listing.
		if !team.IsNational || t.players[player.ID] == nil {
			t.players[player.ID] = player
		}
	}
}

func (t *targets) IsComplete() bool {
	for _, team := range t.teams {
		if team == nil {
			return false
		}
	}
	return true
}

func (t *targets) PlayersFound() (players teams.Players) {
	for _, player := range t.players {
		players = append(players, player)
	}
	return
}
