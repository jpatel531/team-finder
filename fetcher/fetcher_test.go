package fetcher

import (
	"fmt"
	"github.com/jpatel531/team-finder/teams"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOnStatusOKEndpointScrapedAndTeamReturned(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/teams/1", req.URL.Path)
			w.Write([]byte(fixture))
		}),
	)

	f := New(fmt.Sprintf("%s/%s", server.URL, "teams/%d"))

	team, err := f.Fetch(1)

	assert.NoError(t, err)
	assert.Equal(t, &teams.Team{
		ID:   1,
		Name: "Apoel FC",
		Matches: teams.Matches{
			Last: teams.Match{
				Teamhome: teams.Opponent{ID: 1, Name: "Apoel FC"},
				Teamaway: teams.Opponent{ID: 1874, Name: "FC Astana"}},
			Next: teams.Match{
				Teamhome: teams.Opponent{ID: 24, Name: "Olympiakos"},
				Teamaway: teams.Opponent{ID: 1, Name: "Apoel FC"}},
			Following: teams.Match{
				Teamhome: teams.Opponent{ID: 347, Name: "BSC YB"},
				Teamaway: teams.Opponent{ID: 1, Name: "Apoel FC"}},
		},
		Players: []*teams.Player{
			{ID: "6", Country: "Portugal", Name: "Nuno Morais", Age: "32"},
			{ID: "19", Country: "Cyprus", Name: "Nektarious Alexandrou", Age: "32"},
			{ID: "770", Country: "Spain", Name: "Urko Pardo", Age: "33"},
		},
		IsNational: false,
	}, team)
}

func TestCannotScrapeTheSameTeamMoreThanOnce(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/teams/1", req.URL.Path)
			w.Write([]byte(fixture))
		}),
	)

	f := New(fmt.Sprintf("%s/%s", server.URL, "teams/%d"))

	_, err := f.Fetch(1)
	assert.NoError(t, err)

	_, err = f.Fetch(1)
	assert.Equal(t, AlreadyScraped, err)
}

func Test404ReturnedAsNotFound(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}),
	)

	f := New(fmt.Sprintf("%s/%s", server.URL, "teams/%d"))

	_, err := f.Fetch(32)
	assert.Equal(t, NotFound, err)
}

func TestNon200OrNon400ReturnedAsGenericError(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}),
	)

	f := New(fmt.Sprintf("%s/%s", server.URL, "teams/%d"))

	_, err := f.Fetch(90)
	assert.Equal(t, RequestError, err)
}

func TestNilTeamReturnedAsError(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(`
				{
					"status": "ok",
					"code": 0,
					"data": {},
					"message": "woopsie"
				}
			`))
		}),
	)
	f := New(fmt.Sprintf("%s/%s", server.URL, "teams/%d"))

	_, err := f.Fetch(40)
	assert.Equal(t, NilTeam, err)
}
