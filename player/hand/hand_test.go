package hand

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card/equipment"
	"github.com/hatajoe/hatastone/player/card/minion"
)

func TestPopByIndex(t *testing.T) {
	h := NewHand()
	h.Push(minion.NewMurloc())
	h.Push(equipment.NewWeapon())
	if actual, ok := h.PopByIndex(1).(*equipment.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &equipment.Weapon{}, actual)
	}
	if c := h.PopByIndex(-1); c != nil {
		t.Fatalf("unexpected value popped. expected=nil, actual=%T", c)
	}
}
