package match

import (
	"testing"
	"time"

	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/proxy"
	"github.com/hatajoe/hatastone/match/rule"
	"github.com/hatajoe/hatastone/player/card"
	player "github.com/hatajoe/hatastone/player/context/match"
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestExec(t *testing.T) {
	p1 := player.NewPlayer(
		"p1",
		deck.NewDeck(hero.NewMage(20), []card.ICard{
			card.NewMurloc("p1m1", 2, 2),
			card.NewWeapon("p1w1", 2),
			card.NewInstant("p1i1", 5),
			card.NewMurloc("p1m2", 2, 2),
			card.NewWeapon("p1w2", 2),
			card.NewInstant("p1i2", 5),
		}),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)
	p2 := player.NewPlayer(
		"p2",
		deck.NewDeck(hero.NewMage(20), []card.ICard{
			card.NewMurloc("p1m1", 2, 2),
			card.NewWeapon("p1w1", 2),
			card.NewInstant("p1i1", 5),
			card.NewMurloc("p1m2", 2, 2),
			card.NewWeapon("p1w2", 2),
			card.NewInstant("p1i2", 5),
		}),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)

	p1CoinTossCh := proxy.NewCoinTossProxy(p1).Listen().(event.CoinToss)
	defer close(p1CoinTossCh)
	p2CoinTossCh := proxy.NewCoinTossProxy(p2).Listen().(event.CoinToss)
	defer close(p2CoinTossCh)
	p1DrawCh := proxy.NewDrawProxy(p1).Listen().(event.Draw)
	defer close(p1DrawCh)
	p2DrawCh := proxy.NewDrawProxy(p2).Listen().(event.Draw)
	defer close(p2DrawCh)
	p1MariganCh := proxy.NewMariganProxy(p1).Listen().(event.Marigan)
	defer close(p1MariganCh)
	p2MariganCh := proxy.NewMariganProxy(p2).Listen().(event.Marigan)
	defer close(p2MariganCh)

	ctx := NewContext(
		rule.NewRule(
			time.Now().UnixNano(),
			3,
		),
		event.Events{
			p1CoinTossCh,
			p1DrawCh,
			p1MariganCh,
		},
		event.Events{
			p2CoinTossCh,
			p2DrawCh,
			p2MariganCh,
		},
	)

	for ctx.GetState() != nil {
		if err := ctx.Exec(); err != nil {
			t.Fatal(err)
		}
	}
}
