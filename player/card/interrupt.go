package card

type Interrupt struct {
	Entity
}

func NewInterrupt(id string) *Interrupt {
	return &Interrupt{
		Entity: Entity{
			ID: id,
		},
	}
}
