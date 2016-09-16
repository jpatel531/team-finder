package targets

import (
	"github.com/jpatel531/team-finder/teams"
	"log"
	"sync"
)

/*
Targets is an interface that tracks the progress of our search.

It holds in its underlying data the teams we are looking for and, if they are found,
the data associated with it. It also holds those teams' players which we will need to
render once the search is complete.

It holds a sync.RWMutex to ensure safe concurrent access to its underlying data.
*/
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

// A team is only targeted if its name is a key in the underlying
// teams map, and if the value of that key is nil.
// Because when a target is hit, the value becomes non-nil, a team
// that has been found is no longer targeted.
func (t *targets) TeamIsTargeted(teamName string) bool {
	t.RLock()
	defer t.RUnlock()
	team, exists := t.teams[teamName]
	return exists && team == nil
}

/*
One calls `Hit` when a target has been hit, i.e. when a team has been found.

We store it in the underlying `teams` map, and its players in our underlying `players`
map.
*/
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

// All the targets have been found if all its teams are not nil.
func (t *targets) IsComplete() bool {
	for _, team := range t.teams {
		if team == nil {
			return false
		}
	}
	return true
}

// Returns all the targeted players
func (t *targets) PlayersFound() (players teams.Players) {
	for _, player := range t.players {
		players = append(players, player)
	}
	return
}
