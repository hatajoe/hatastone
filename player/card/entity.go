package card

type Entity struct {
	id string
}

func NewEntity(id string) *Entity {
	return &Entity{
		id: id,
	}
}

func (e Entity) GetID() string {
	return e.id
}
