package match

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/card/equipment"
	"github.com/hatajoe/hatastone/player/card/minion"
	"github.com/hatajoe/hatastone/player/card/spel"
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
				minion.NewMurloc(),
				equipment.NewWeapon(),
				spel.NewInstant(),
			},
		),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)

	if !p.Draw() {
		t.Fatalf("Draw() is failed. deck is empty")
	}
	if err := p.Marigan([]int{0, 2}); err != nil {
		t.Fatalf("Marigan() is failed. %s", err)
	}
}
