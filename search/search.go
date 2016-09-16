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

// The private search struct available only in this package.
// It is composed of targets and an HTTP fetcher.
type search struct {
	targets.Targets
	fetcher.Fetcher
}

/*
DoBatch iterates through a given range of numbers.

For each number it tries to send a request for a team with that number as its ID. For
this request, it spawns a new goroutine.

If that request results in a 404, it increments an atomic counter. If this atomic counter
exceeds the threshold for the batch, it returns false, assuming that we have exceeded
the number of teams there are in the API.

If it has not exceeded this threshold, and the search is still incomplete, it calls DoBatch
on the next range.

If the search is complete, it returns true.
*/
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

// For a given ID, a request is initiated for its contents. If its name is one of our targets
// we store it in the targets.Targets structure.
//
// If the search is unsuccessful, we scan the team's previous, next and following opponents
// to see if they're a team target. If so, we follow its ID to retrieve the team data.
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
