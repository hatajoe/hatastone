package hero

import "github.com/hatajoe/hatastone/player/card"

type Hero struct {
	Life      int
	Equipment card.IEquipment
}

func NewHero(life int) *Hero {
	return &Hero{
		Life: life,
	}
}

func (h Hero) IsDead() bool {
	return h.Life <= 0
}

func (h Hero) GetLife() int {
	return h.Life
}

func (h Hero) GetEquipment() card.IEquipment {
	return h.Equipment
}

func (h *Hero) AttachEquipment(e card.IEquipment) {
	h.Equipment = e
}

func (h *Hero) DetachEquipment() {
	h.Equipment = nil
}

func (h Hero) GetAttack() int {
	if h.Equipment == nil {
		return 0
	}
	return h.Equipment.GetAttack()
}

func (h *Hero) AcceptAttack(damage int) {
	h.Life -= damage
}
