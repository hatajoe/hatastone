package hand

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAdd(t *testing.T) {
	expected := card.Cards{
		card.NewMurloc(card.NewEntity("murloc")),
		card.NewWeapon(card.NewEntity("weapon")),
		card.NewInstant(card.NewEntity("instant")),
	}

	h := NewHand()
	h.Add(card.NewMurloc(card.NewEntity("murloc")))
	h.Add(card.NewWeapon(card.NewEntity("weapon")))
	h.Add(card.NewInstant(card.NewEntity("instant")))

	for i, card := range h.GetCards() {
		if expected[i].GetID() != card.GetID() {
			t.Errorf("field.Add is unexpected behavior. expected=%s, actual=%s", expected[i].GetID(), card.GetID())
		}
	}
}

func TestRemoveByID(t *testing.T) {
	h := NewHand()
	h.Add(card.NewMurloc(card.NewEntity("murloc")))
	h.Add(card.NewWeapon(card.NewEntity("weapon")))
	if actual, ok := h.RemoveByID("weapon").(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
	if c := h.RemoveByID("weapon"); c != nil {
		t.Fatalf("unexpected value popped. expected=nil, actual=%T", c)
	}
	if c := h.RemoveByID("pirates"); c != nil {
		t.Fatalf("unexpected value popped. expected=nil, actual=%T", c)
	}
}
