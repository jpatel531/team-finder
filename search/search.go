package search

import (
	"github.com/jpatel531/team-finder/fetcher"
	"github.com/jpatel531/team-finder/targets"
	"github.com/jpatel531/team-finder/teams"
	"sync"
	"sync/atomic"
)

type Search interface {
	DoBatch(startID, size, incr, notFoundThreshold int) (complete bool)
	PlayersFound() teams.Players
}

func New(endpoint string, teamNames []string) Search {
	return &search{
		Targets: targets.New(teamNames),
		Fetcher: fetcher.New(endpoint),
	}
}

type search struct {
	targets.Targets
	fetcher.Fetcher
}

func (s *search) DoBatch(startID, size, incr, notFoundThreshold int) (complete bool) {
	var (
		wg            sync.WaitGroup
		i             int
		notFoundCount int32
	)

	for i = startID; i < startID+size; i += incr {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := s.tryTeam(id)
			if err == fetcher.NotFound {
				atomic.AddInt32(&notFoundCount, 1)
			}
		}(i)
	}
	wg.Wait()

	if atomic.LoadInt32(&notFoundCount) >= int32(notFoundThreshold) {
		return false
	}

	if !s.IsComplete() {
		return s.DoBatch(i, size, incr, notFoundThreshold)
	}
	return true
}

func (s *search) tryTeam(teamID int) (err error) {
	if s.IsComplete() {
		return
	}
	var team *teams.Team

	team, err = s.Fetch(teamID)
	if err != nil {
		return
	}
	name := team.Name
	found := s.TeamIsTargeted(name)

	if found {
		s.Hit(team)
	} else {
		// are any of its opponents targeted by us?
		s.tryOpponents(team)
	}
	return
}

func (s *search) tryOpponents(team *teams.Team) {
	opponents := team.Opponents()
	for _, opponent := range opponents {
		if s.TeamIsTargeted(opponent.Name) {
			team, err := s.Fetch(opponent.ID)
			if err != nil {
				return
			}
			s.Hit(team)
		}
	}
}
