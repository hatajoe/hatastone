package hand

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestRemove(t *testing.T) {
	h := NewHand()
	h.Add(card.NewMurloc(card.NewEntity("murloc")))
	h.Add(card.NewWeapon(card.NewEntity("weapon")))
	if actual, ok := h.RemoveByID("weapon").(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
	if c := h.RemoveByID("pirates"); c != nil {
		t.Fatalf("unexpected value popped. expected=nil, actual=%T", c)
	}
}
