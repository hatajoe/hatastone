package card

type Interrupt struct {
	entity
}

func NewInterrupt(id string) *Interrupt {
	return &Interrupt{
		entity: entity{
			id: id,
		},
	}
}
