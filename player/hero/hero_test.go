package hero

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAttachEquipment(t *testing.T) {
	h := NewHero()
	h.AttachEquipment(card.NewWeapon(card.NewEntity("weapon")))
	if actual, ok := h.GetEquipment().(*card.Weapon); !ok {
		t.Fatalf("unexpected equipment was found. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}

func TestDetachEquipment(t *testing.T) {
	h := NewHero()
	h.AttachEquipment(card.NewWeapon(card.NewEntity("weapon")))
	h.DetachEquipment()
	if h.GetEquipment() != nil {
		t.Fatalf("DetachEquipment failed. expected=nil, actual=%T", h.GetEquipment())
	}
}
