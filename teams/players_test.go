package teams

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortingPlayersSortsThemAlphabetically(t *testing.T) {
	players := Players([]*Player{
		{Name: "Zinedine Zidane"},
		{Name: "Robin Van Persie"},
		{Name: "Artur Boruc"},
		{Name: "David Luiz"},
		{Name: "Marco Reus"},
		{Name: "Mario Götze"},
	})
	sort.Sort(players)
	assert.Equal(t, Players([]*Player{
		{Name: "Artur Boruc"},
		{Name: "David Luiz"},
		{Name: "Marco Reus"},
		{Name: "Mario Götze"},
		{Name: "Robin Van Persie"},
		{Name: "Zinedine Zidane"},
	}), players)
}

func TestStringifyingPlayersFormatsThemCorrectly(t *testing.T) {
	players := Players([]*Player{
		{Name: "Alexander Mustermann", Age: "25", Country: "France", Club: "Manchester Utd"},
		{Name: "Brad Examplemann", Age: "30", Country: "Switzerland"},
	})
	assert.Equal(t, `Alexander Mustermann; 25; France, Manchester Utd
Brad Examplemann; 30; Switzerland
`, players.String())
}
