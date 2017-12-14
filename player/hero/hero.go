package hero

import "github.com/hatajoe/hatastone/player/card"

type Hero struct {
	life      int
	equipment card.IEquipment
}

func NewHero(life int) *Hero {
	return &Hero{
		life: life,
	}
}

func (h Hero) GetLife() int {
	return h.life
}

func (h Hero) GetEquipment() card.IEquipment {
	return h.equipment
}

func (h *Hero) AttachEquipment(e card.IEquipment) {
	h.equipment = e
}

func (h *Hero) DetachEquipment() {
	h.equipment = nil
}

func (h Hero) GetAttack() int {
	if h.equipment == nil {
		return 0
	}
	return h.equipment.GetAttack()
}

func (h *Hero) AcceptAttack(damage int) {
	h.life -= damage
}
