package hero

import "github.com/hatajoe/hatastone/player/card"

type Hero struct {
	equipment card.IEquipment
}

func NewHero() *Hero {
	return &Hero{}
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
