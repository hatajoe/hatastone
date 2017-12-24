package card

type Minion struct {
	Atk  int
	Life int
}

func (m Minion) GetLife() int {
	return m.Life
}

func (m Minion) GetAttack() int {
	return m.Atk
}

func (m *Minion) AcceptAttack(damage int) {
	m.Life -= damage
	if m.Life < 0 {
		m.Life = 0
	}
}
