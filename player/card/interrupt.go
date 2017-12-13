package card

type Interrupt struct {
	*Entity
}

func NewInterrupt(e *Entity) *Interrupt {
	return &Interrupt{Entity: e}
}
