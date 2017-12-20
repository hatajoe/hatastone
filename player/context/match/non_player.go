package match

type NonPlayer struct {
	*match
}

func NewNonPlayer(m *match) *NonPlayer {
	return &NonPlayer{match: m}
}
