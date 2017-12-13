package card

type Instant struct {
	*Entity
}

func NewInstant(e *Entity) *Instant {
	return &Instant{Entity: e}
}
