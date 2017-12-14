package hero

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAttachEquipment(t *testing.T) {
	h := NewHero(20)
	h.AttachEquipment(card.NewWeapon("weapon", 2))
	if actual, ok := h.GetEquipment().(*card.Weapon); !ok {
		t.Fatalf("unexpected equipment was found. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}

func TestDetachEquipment(t *testing.T) {
	h := NewHero(20)
	h.AttachEquipment(card.NewWeapon("weapon", 2))
	h.DetachEquipment()
	if h.GetEquipment() != nil {
		t.Fatalf("DetachEquipment failed. expected=nil, actual=%T", h.GetEquipment())
	}
}
