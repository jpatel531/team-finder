package targets

import (
	"github.com/jpatel531/team-finder/teams"
	"github.com/stretchr/testify/assert"
	"sort"
	"sync"
	"testing"
)

func TestTeamIsTargetedIsTrueIfTeamInMapAndNotNil(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": nil,
	}
	tgs := newTestTargets(clubs, nil)
	assert.True(t, tgs.TeamIsTargeted("Chelsea"))
}

func TestTeamTargetedIfNoTeamInMap(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": nil,
	}
	tgs := newTestTargets(clubs, nil)
	assert.False(t, tgs.TeamIsTargeted("Southamptom"))
}

func TestTeamTargetedIsFalseIfNotNil(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": &teams.Team{},
	}
	tgs := newTestTargets(clubs, nil)
	assert.False(t, tgs.TeamIsTargeted("Chelsea"))
}

func TestHitPanicsIfTeamIsntTargeted(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": &teams.Team{},
	}
	tgs := newTestTargets(clubs, nil)
	assert.Panics(t, func() {
		tgs.Hit(&teams.Team{Name: "Southampton"})
	})
}

func TestHitStoresTeamAndPlayersInTeamsMap(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": nil,
	}
	tgs := newTestTargets(clubs, make(map[string]*teams.Player))
	found := &teams.Team{
		Name: "Chelsea",
		Players: []*teams.Player{
			{ID: "4", Country: "Spain", Name: "Chicken Cesar Azpilicueta"},
			{ID: "9", Country: "France", Name: "Ngolo Kanté"},
		},
	}
	tgs.Hit(found)
	assert.Equal(t, found, tgs.teams["Chelsea"])
	assert.Equal(t, map[string]*teams.Player{
		"4": {Club: "Chelsea", ID: "4", Country: "Spain", Name: "Chicken Cesar Azpilicueta"},
		"9": {Club: "Chelsea", ID: "9", Country: "France", Name: "Ngolo Kanté"},
	}, tgs.players)
}

func TestHitWithClubHasPriorityOverHitWithInternationalClub(t *testing.T) {
	clubs := map[string]*teams.Team{
		"Chelsea": nil,
		"Spain":   nil,
		"Belgium": nil,
	}
	tgs := newTestTargets(clubs, make(map[string]*teams.Player))

	azpiSpain := &teams.Player{ID: "4", Country: "Spain", Name: "Chicken Cesar Azpilicueta"}

	spain := &teams.Team{
		Name:       "Spain",
		Players:    []*teams.Player{azpiSpain},
		IsNational: true,
	}

	tgs.Hit(spain)
	assert.Equal(t, azpiSpain, tgs.players["4"])

	// testing club overrides country
	azpiChelsea := &teams.Player{ID: "4", Country: "Spain", Name: "Chicken Cesar Azpilicueta"}
	hazardChelsea := &teams.Player{ID: "18", Country: "Belgium", Name: "Eden Hazard"}

	chelsea := &teams.Team{
		Name:       "Chelsea",
		Players:    []*teams.Player{azpiChelsea, hazardChelsea},
		IsNational: false,
	}

	tgs.Hit(chelsea)
	assert.Equal(t, azpiChelsea, tgs.players["4"])

	// testing country doesn't override club
	belgium := &teams.Team{
		Name:       "Belgium",
		Players:    []*teams.Player{{ID: "18", Country: "Belgium", Name: "Eden Hazard"}},
		IsNational: true,
	}
	tgs.Hit(belgium)
	assert.Equal(t, hazardChelsea, tgs.players["18"])
}

func TestIsCompleteIfAllTeamsAreNotNull(t *testing.T) {
	tgs := newTestTargets(map[string]*teams.Team{
		"Arsenal":         nil,
		"Manchester City": nil,
	}, nil)
	assert.False(t, tgs.IsComplete())

	tgs.Hit(&teams.Team{Name: "Arsenal"})
	assert.False(t, tgs.IsComplete())

	tgs.Hit(&teams.Team{Name: "Manchester City"})
	assert.True(t, tgs.IsComplete())
}

func TestPlayersReturnsAllPlayersInMapAsPlayersType(t *testing.T) {
	hazard := &teams.Player{Name: "Eden Hazard"}
	rom := &teams.Player{Name: "Romelu Lukaku"}

	tgs := newTestTargets(nil, map[string]*teams.Player{
		"18": hazard,
		"19": rom,
	})
	expected := teams.Players([]*teams.Player{hazard, rom})
	sort.Sort(expected) // avoid ordering issues

	actual := tgs.PlayersFound()
	sort.Sort(actual)
	assert.Equal(t, expected, actual)
}

func newTestTargets(teams map[string]*teams.Team, players map[string]*teams.Player) *targets {
	return &targets{
		RWMutex: &sync.RWMutex{},
		teams:   teams,
		players: players,
	}
}
