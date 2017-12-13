package card

type Weapon struct {
	*Entity
}

func NewWeapon(e *Entity) *Weapon {
	return &Weapon{Entity: e}
}
