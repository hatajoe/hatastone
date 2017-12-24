package card

type Spel struct {
	Atk int
}

func (s Spel) GetAttack() int {
	return s.Atk
}
