package card

type Entity struct {
	ID string
}

func (e Entity) GetID() string {
	return e.ID
}
