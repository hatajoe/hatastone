package hand

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestPopByID(t *testing.T) {
	h := NewHand()
	h.Push(card.NewMurloc(card.NewEntity("murloc")))
	h.Push(card.NewWeapon(card.NewEntity("weapon")))
	if actual, ok := h.PopByID("weapon").(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
	if c := h.PopByID("pirates"); c != nil {
		t.Fatalf("unexpected value popped. expected=nil, actual=%T", c)
	}
}
