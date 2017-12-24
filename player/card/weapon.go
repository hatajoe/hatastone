package card

type Weapon struct {
	Entity
	Atk int
}

func NewWeapon(id string, atk int) *Weapon {
	return &Weapon{
		Entity: Entity{
			ID: id,
		},
		Atk: atk,
	}
}

func (w Weapon) GetAttack() int {
	return w.Atk
}
