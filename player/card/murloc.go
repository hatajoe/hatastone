package card

type Murloc struct {
	*Entity
}

func NewMurloc(e *Entity) *Murloc {
	return &Murloc{Entity: e}
}
