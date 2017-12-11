package match

type Player struct {
	*Match
}

func NewPlayer(m *Match) *Player {
	return &Player{Match: m}
}
