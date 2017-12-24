package card

type Murloc struct {
	Entity
	Minion
}

func NewMurloc(id string, atk int, life int) *Murloc {
	return &Murloc{
		Entity: Entity{
			ID: id,
		},
		Minion: Minion{
			Atk:  atk,
			Life: life,
		},
	}
}
