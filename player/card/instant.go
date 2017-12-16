package card

type Instant struct {
	entity
	spel
}

func NewInstant(id string, atk int) *Instant {
	return &Instant{
		entity: entity{
			id: id,
		},
		spel: spel{
			atk: atk,
		},
	}
}
