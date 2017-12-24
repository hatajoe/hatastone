package card

type Instant struct {
	Entity
	Spel
}

func NewInstant(id string, atk int) *Instant {
	return &Instant{
		Entity: Entity{
			ID: id,
		},
		Spel: Spel{
			Atk: atk,
		},
	}
}
