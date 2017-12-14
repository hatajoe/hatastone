package card

type entity struct {
	id string
}

func (e entity) GetID() string {
	return e.id
}
