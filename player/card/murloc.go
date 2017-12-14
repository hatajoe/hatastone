package card

type Murloc struct {
	entity
	minion
}

func NewMurloc(id string, atk int, life int) *Murloc {
	return &Murloc{
		entity: entity{
			id: id,
		},
		minion: minion{
			atk:  atk,
			life: life,
		},
	}
}
