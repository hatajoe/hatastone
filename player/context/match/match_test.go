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
	p := NewMatch(
		deck.NewDeck(
			hero.NewMage(),
			[]card.ICard{
				card.NewMurloc(card.NewEntity("murloc")),
				card.NewWeapon(card.NewEntity("weapon")),
				card.NewInstant(card.NewEntity("instant")),
			},
		),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)

	if !p.Draw() {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if !p.Draw() {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if !p.Draw() {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if err := p.Marigan([]string{"murloc", "instant"}); err != nil {
		t.Fatalf("Marigan() is failed. %s", err)
	}
}
