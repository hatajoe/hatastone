package card

type Cards []ICard

func (m Cards) FindByID(id string) ICard {
	for _, c := range m {
		if c.GetID() == id {
			return c
		}
	}
	return nil
}

func (m *Cards) DeleteByID(id string) {
	for i, c := range *m {
		if c.GetID() == id {
			*m = append((*m)[:i], (*m)[i+1:]...)
			break
		}
	}
}
