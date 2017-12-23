package match

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
	"github.com/hatajoe/hatastone/player/hero"
)

func Test(t *testing.T) {
	p := NewPlayer(
		"player1",
		deck.NewDeck(
			hero.NewMage(20),
			[]card.ICard{
				card.NewMurloc("murloc", 2, 2),
				card.NewWeapon("weapon", 2),
				card.NewInstant("instant", 4),
			},
		),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)

	if c := p.Draw(); c == nil {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if c := p.Draw(); c == nil {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if c := p.Draw(); c == nil {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if c := p.Marigan([]string{"murloc", "instant"}); len(c) <= 0 {
		t.Fatalf("Marigan() is failed. deck is empty")
	}
	if err := p.Play("murloc", 0); err != nil {
		t.Fatalf("Play() is failed. %s", err)
	}
	if err := p.DiscardFromHand("weapon"); err != nil {
		t.Fatalf("DiscardFromHand() is failed. %s", err)
	}
	if err := p.DiscardFromField("murloc"); err != nil {
		t.Fatalf("DiscardFromField() is failed. %s", err)
	}
}
