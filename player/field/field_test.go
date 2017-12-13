package field

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAdd(t *testing.T) {
	expected := card.Cards{
		card.NewMurloc(card.NewEntity("murloc")),
		card.NewInstant(card.NewEntity("instant")),
		card.NewWeapon(card.NewEntity("weapon")),
	}

	h := NewField()
	h.Add(card.NewMurloc(card.NewEntity("murloc")), 0)
	h.Add(card.NewWeapon(card.NewEntity("weapon")), 1)
	h.Add(card.NewInstant(card.NewEntity("instant")), 1)

	for i, card := range h.GetCards() {
		if expected[i].GetID() != card.GetID() {
			t.Errorf("field.Add is unexpected behavior. expected=%s, actual=%s", expected[i].GetID(), card.GetID())
		}
	}
}

func TestRemoveByID(t *testing.T) {
	h := NewField()
	if err := h.Add(card.NewWeapon(card.NewEntity("weapon")), 1); err == nil {
		t.Fatalf("error shuould be occured")
	}
	if err := h.Add(card.NewWeapon(card.NewEntity("weapon")), 0); err != nil {
		t.Fatalf("unexpected error occured. %s", err)
	}
	if actual, ok := h.RemoveByID("weapon").(*card.Weapon); !ok {
		t.Fatalf("unexpected card is removed. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
	if c := h.RemoveByID("weapon"); c != nil {
		t.Fatalf("unexpected value removed. expected=nil, actual=%T", c)
	}
	if c := h.RemoveByID("pirates"); c != nil {
		t.Fatalf("unexpected value removed. expected=nil, actual=%T", c)
	}
}
