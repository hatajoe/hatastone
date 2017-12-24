package match

import (
	"testing"
	"time"

	"github.com/hatajoe/hatastone/apps/console"
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
			card.NewMurloc("m1", 2, 2),
			card.NewWeapon("w1", 2),
			card.NewInstant("i1", 5),
			card.NewMurloc("m2", 2, 2),
			card.NewWeapon("w2", 2),
			card.NewInstant("i2", 5),
		}),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)
	p2 := player.NewPlayer(
		"p2",
		deck.NewDeck(hero.NewMage(20), []card.ICard{
			card.NewMurloc("m1", 2, 2),
			card.NewWeapon("w1", 2),
			card.NewInstant("i1", 5),
			card.NewMurloc("m2", 2, 2),
			card.NewWeapon("w2", 2),
			card.NewInstant("i2", 5),
		}),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)
	r := &console.Reader{}
	w := &console.Writer{}

	p1CoinTossCh := proxy.NewCoinTossProxy(p1).Listen(r, w).(event.CoinToss)
	defer close(p1CoinTossCh)
	p2CoinTossCh := proxy.NewCoinTossProxy(p2).Listen(r, w).(event.CoinToss)
	defer close(p2CoinTossCh)
	p1DrawCh := proxy.NewDrawProxy(p1).Listen(r, w).(event.Draw)
	defer close(p1DrawCh)
	p2DrawCh := proxy.NewDrawProxy(p2).Listen(r, w).(event.Draw)
	defer close(p2DrawCh)
	p1MariganCh := proxy.NewMariganProxy(p1).Listen(r, w).(event.Marigan)
	defer close(p1MariganCh)
	p2MariganCh := proxy.NewMariganProxy(p2).Listen(r, w).(event.Marigan)
	defer close(p2MariganCh)
	p1PlayCh := proxy.NewPlayProxy(p1).Listen(r, w).(event.Play)
	defer close(p1PlayCh)
	p2PlayCh := proxy.NewPlayProxy(p2).Listen(r, w).(event.Play)
	defer close(p2PlayCh)

	ctx := NewContext(
		rule.NewRule(
			time.Now().UnixNano(),
			3,
		),
		p1,
		p2,
		event.Events{
			p1CoinTossCh,
			p1DrawCh,
			p1MariganCh,
			p1PlayCh,
		},
		event.Events{
			p2CoinTossCh,
			p2DrawCh,
			p2MariganCh,
			p2PlayCh,
		},
	)

	for ctx.GetState() != nil {
		if err := ctx.Exec(); err != nil {
			t.Fatal(err)
		}
	}
}
