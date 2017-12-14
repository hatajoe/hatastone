package card

type Instant struct {
	entity
}

func NewInstant(id string) *Instant {
	return &Instant{
		entity: entity{
			id: id,
		},
	}
}
