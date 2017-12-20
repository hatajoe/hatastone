package discard

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAdd(t *testing.T) {
	expected := card.Cards{
		card.NewMurloc("murloc", 2, 2),
		card.NewInstant("instant", 2),
		card.NewWeapon("weapon", 2),
	}

	h := NewDiscard()
	h.Add(card.NewWeapon("weapon", 2))
	h.Add(card.NewInstant("instant", 2))
	h.Add(card.NewMurloc("murloc", 2, 2))

	for i, card := range h.GetCards() {
		if expected[i].GetID() != card.GetID() {
			t.Errorf("discard.Add is unexpected behavior. expected=%s, actual=%s", expected[i].GetID(), card.GetID())
		}
	}
}
