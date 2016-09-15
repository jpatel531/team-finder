package fetcher

import "github.com/jpatel531/team-finder/teams"

type response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team *teams.Team `json:"team"`
	} `json:"data"`
}

func (r response) team() *teams.Team {
	return r.Data.Team
}
