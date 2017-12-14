package card

type minion struct {
	atk  int
	life int
}

func (m minion) GetLife() int {
	return m.life
}

func (m minion) GetAttack() int {
	return m.atk
}

func (m *minion) AcceptAttack(damage int) {
	m.life -= damage
	if m.life < 0 {
		m.life = 0
	}
}
