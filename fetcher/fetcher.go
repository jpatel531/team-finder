package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jpatel531/team-finder/teams"
	"io/ioutil"
	"net/http"
)

// A Fetcher's purpose is to fetch the team with a given ID at an endpoint,
// parse it, and return it, as well as any errors encountered.
// It does not duplicate work. It holds a set of integer IDs and checks
// that the ID has not been previously scraped before doing its work.
type Fetcher interface {
	Fetch(id int) (team *teams.Team, err error)
}

func New(endpoint string) Fetcher {
	return &fetcher{
		endpoint:       endpoint,
		alreadyFetched: newIntSet(),
	}
}

type fetcher struct {
	endpoint       string
	alreadyFetched *intSet
}

var (
	NotFound       = errors.New("NotFound")
	RequestError   = errors.New("RequestError")
	NilTeam        = errors.New("NilTeam")
	AlreadyScraped = errors.New("AlreadyScraped")
)

func (f *fetcher) Fetch(id int) (team *teams.Team, err error) {
	var parsed *response
	if f.alreadyFetched.exists(id) {
		err = AlreadyScraped
		return
	}
	f.alreadyFetched.add(id)

	resp, err := http.Get(fmt.Sprintf(f.endpoint, id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			err = NotFound
		} else {
			// we're not too worried about other types of status codes.
			// We MAY want to retry 5xx, but let's keep this exercise simple
			err = RequestError
		}
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &parsed); err != nil {
		return
	}
	team = parsed.team()
	if team == nil {
		err = NilTeam
	}
	return
}
