package card

type spel struct {
	atk int
}

func (s spel) GetAttack() int {
	return s.atk
}
