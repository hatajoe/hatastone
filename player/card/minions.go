package card

type Minions []IMinion

func (m Minions) FindByID(id string) IMinion {
	for _, c := range m {
		if c.GetID() == id {
			return c
		}
	}
	return nil
}

func (m *Minions) DeleteByID(id string) {
	for i, c := range *m {
		if c.GetID() == id {
			*m = append((*m)[:i], (*m)[i+1:]...)
			break
		}
	}
}
