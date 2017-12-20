package match

import (
	"testing"
	"time"

	"github.com/hatajoe/hatastone/match/event"
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
		}),
		hand.NewHand(),
		field.NewField(),
		discard.NewDiscard(),
	)

	p1CoinTossCh := make(event.CoinToss)
	defer close(p1CoinTossCh)
	p2CoinTossCh := make(event.CoinToss)
	defer close(p2CoinTossCh)
	p1DrawCh := make(event.Draw)
	defer close(p1DrawCh)
	p2DrawCh := make(event.Draw)
	defer close(p2DrawCh)

	ctx := NewContext(
		rule.NewRule(
			time.Now().UnixNano(),
			3,
		),
		event.Events{
			p1DrawCh,
			p1CoinTossCh,
		},
		event.Events{
			p2DrawCh,
			p2CoinTossCh,
		},
	)

	go func() {
		for {
			select {
			case order, ok := <-p1CoinTossCh:
				if ok {
					t.Logf("p1 order is %d", order)
				}
			case order, ok := <-p2CoinTossCh:
				if ok {
					t.Logf("p2 order is %d", order)
				}
			case _, ok := <-p1DrawCh:
				if ok {
					c := p1.Draw()
					if c == nil {
						t.Fatal("p1 deck is empty")
					}
					t.Logf("p1 draw %s", c.GetID())
				}
			case _, ok := <-p2DrawCh:
				if ok {
					c := p2.Draw()
					if c == nil {
						t.Fatal("p2 deck is empty")
					}
					t.Logf("p2 draw %s", c.GetID())
				}
			default:
			}
		}
	}()

	for ctx.GetState() != nil {
		if err := ctx.Exec(); err != nil {
			t.Fatal(err)
		}
	}
}
