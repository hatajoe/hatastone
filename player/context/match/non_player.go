package match

type NonPlayer struct {
	*Match
}

func NewNonPlayer(m *Match) *NonPlayer {
	return &NonPlayer{Match: m}
}
