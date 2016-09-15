package search

import (
	"github.com/jpatel531/team-finder/fetcher"
	"github.com/jpatel531/team-finder/mocks"
	"github.com/jpatel531/team-finder/teams"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTryTeamHitsTargetIfTeamFound(t *testing.T) {
	f := &mocks.Fetcher{}

	team := &teams.Team{ID: 88, Name: "Wigan Athletic"}
	f.On("Fetch", 94).Return(team, nil)

	tgs := &mocks.Targets{}
	tgs.
		On("IsComplete").
		Return(false).
		On("TeamIsTargeted", "Wigan Athletic").
		Return(true).
		On("Hit", team)

	s := &search{tgs, f}

	err := s.tryTeam(94)
	assert.NoError(t, err)
	tgs.AssertExpectations(t)
	f.AssertExpectations(t)
}

func TestTryTeamTriesOpponentsIfNotFound(t *testing.T) {
	f := &mocks.Fetcher{}

	// from the list of opponents, the next hop will be to Crawley Town
	wigan := &teams.Team{
		ID:   88,
		Name: "Wigan Athletic",
		Matches: teams.Matches{
			Last: teams.Match{
				Teamhome: teams.Opponent{Name: "Wigan Athletic", ID: 88},
				Teamaway: teams.Opponent{Name: "Charlton Athletic", ID: 99},
			},
			Next: teams.Match{
				Teamhome: teams.Opponent{Name: "Wigan Athletic", ID: 88},
				Teamaway: teams.Opponent{Name: "Blackburn Rovers", ID: 104},
			},
			Following: teams.Match{
				Teamhome: teams.Opponent{Name: "Wigan Athletic", ID: 88},
				Teamaway: teams.Opponent{Name: "Crawley Town", ID: 459},
			},
		},
	}

	crawley := &teams.Team{
		ID:   459,
		Name: "Crawley Town",
	}

	f.On("Fetch", 94).Return(wigan, nil).
		On("Fetch", 459).Return(crawley, nil)

	tgs := &mocks.Targets{}
	tgs.
		On("IsComplete").Return(false).
		On("TeamIsTargeted", "Wigan Athletic").Return(false).
		On("TeamIsTargeted", "Charlton Athletic").Return(false).
		On("TeamIsTargeted", "Blackburn Rovers").Return(false).
		On("TeamIsTargeted", "Crawley Town").Return(true). // now pursue Crawley Town!
		On("Hit", crawley)

	s := &search{tgs, f}

	err := s.tryTeam(94)
	assert.NoError(t, err)
	tgs.AssertExpectations(t)
	f.AssertExpectations(t)
}

func TestDoBatchReturnsFalseIfNotFoundCountExceedsThreshold(t *testing.T) {
	f := &mocks.Fetcher{}
	f.
		On("Fetch", 1).Return(nil, fetcher.NotFound).
		On("Fetch", 2).Return(nil, fetcher.NotFound).
		On("Fetch", 3).Return(nil, fetcher.NotFound)

	tgs := &mocks.Targets{}

	tgs.On("IsComplete").Return(false)

	s := &search{tgs, f}

	found := s.DoBatch(1, 3, 1)
	assert.False(t, found)
	tgs.AssertExpectations(t)
	f.AssertExpectations(t)
}

func TestDoBatchTriesWithSubsequentBatchIfNotYetComplete(t *testing.T) {
	startID := 1
	batchSize := 3
	increment := 1

	f := &mocks.Fetcher{}
	tgs := &mocks.Targets{}

	tgs.On("IsComplete").Return(false).Times(7).
		On("TeamIsTargeted", mock.AnythingOfType("string")).Return(true).Times(3).
		On("Hit", mock.AnythingOfType("*teams.Team")).Times(3)

	// for the first round: all of the tries succeed
	// for the second round: all are 404s
	for i := startID; i <= 2*batchSize; i += increment {
		m := f.On("Fetch", i)
		if i <= 3 {
			m.Return(&teams.Team{}, nil)
		} else {
			m.Return(nil, fetcher.NotFound)
		}
	}

	s := &search{tgs, f}
	s.DoBatch(startID, batchSize, increment)
	tgs.AssertExpectations(t)
	f.AssertExpectations(t)
}
