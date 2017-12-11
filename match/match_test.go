package match

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/card/equipment"
	"github.com/hatajoe/hatastone/player/card/minion"
	"github.com/hatajoe/hatastone/player/card/spel"
	pCtxMatch "github.com/hatajoe/hatastone/player/context/match"
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestNewMatch(t *testing.T) {
	p1 := pCtxMatch.NewPlayer(
		pCtxMatch.NewMatch(
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
		),
	)
	p2 := pCtxMatch.NewNonPlayer(
		pCtxMatch.NewMatch(
			deck.NewDeck(
				hero.NewMage(),
				[]card.ICard{
					minion.NewMurloc(),
					equipment.NewWeapon(),
					spel.NewInterrupt(),
				},
			),
			hand.NewHand(),
			field.NewField(),
			discard.NewDiscard(),
		),
	)
	m := NewMatch(p1, p2)
	if err := m.Draw(); err != nil {
		t.Fatal("unexpected deck found")
	}
}
