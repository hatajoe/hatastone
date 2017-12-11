package match

// Match is expressed PVP match
type Match struct {
	player []IPlayer
}

func NewMatch(p1, p2 IPlayer) *Match {
	return &Match{
		player: []IPlayer{p1, p2},
	}
}

func (m *Match) Draw() error {
	for _, p := range m.player {
		if err := p.Draw(); err != nil {
			return err
		}
	}
	return nil
}
