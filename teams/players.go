package teams

import "fmt"

// Each player entry should contain the following information: full name; age; list of teams.
type Player struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Club    string `json:"-"` // an additional field we assign
}

func (p *Player) String() string {
	teams := p.Country
	if p.Club != "" {
		teams = fmt.Sprintf("%s, %s", p.Country, p.Club)
	}
	return fmt.Sprintf("%s; %s; %s\n", p.Name, p.Age, teams)
}

type Players []*Player

// Calling fmt/log.Print on this structure
// will return this return value.
func (p Players) String() (repr string) {
	for _, player := range p {
		repr += player.String()
	}
	return
}

// implementation for the Sort interface. This orders
// players alphabetically by name.
func (p Players) Len() int {
	return len(p)
}

func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Players) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}
