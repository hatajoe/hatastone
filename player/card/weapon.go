package card

type Weapon struct {
	entity
	atk int
}

func NewWeapon(id string, atk int) *Weapon {
	return &Weapon{
		entity: entity{
			id: id,
		},
		atk: atk,
	}
}

func (w Weapon) GetAttack() int {
	return w.atk
}
