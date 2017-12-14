package discard

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

	h := NewDiscard()
	h.Add(card.NewWeapon(card.NewEntity("weapon")))
	h.Add(card.NewInstant(card.NewEntity("instant")))
	h.Add(card.NewMurloc(card.NewEntity("murloc")))

	for i, card := range h.GetCards() {
		if expected[i].GetID() != card.GetID() {
			t.Errorf("discard.Add is unexpected behavior. expected=%s, actual=%s", expected[i].GetID(), card.GetID())
		}
	}
}
