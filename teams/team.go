package teams

type Matches struct {
	Last      Match `json:"last"`
	Next      Match `json:"next"`
	Following Match `json:"following"`
}

type Match struct {
	Teamhome Opponent `json:"teamhome"`
	Teamaway Opponent `json:"teamaway"`
}

type Opponent struct {
	ID   int    `json:"idInternal"`
	Name string `json:"name"`
}

type Team struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Matches    Matches   `json:"matches"`
	Players    []*Player `json:"players"`
	IsNational bool      `json:"isNational"`
}

func (t *Team) Opponents() (opponents []Opponent) {
	matches := t.Matches

	for _, team := range []Opponent{
		matches.Following.Teamaway,
		matches.Following.Teamhome,
		matches.Next.Teamaway,
		matches.Next.Teamhome,
		matches.Last.Teamaway,
		matches.Last.Teamhome,
	} {
		if t.ID != team.ID && team.ID != 0 {
			opponents = append(opponents, team)
		}
	}
	return
}
