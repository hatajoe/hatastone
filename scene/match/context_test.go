package match

import (
	"testing"
	"time"

	"github.com/hatajoe/hatastone/event"
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
		},
		event.Events{
			p2DrawCh,
		},
	)

	go func() {
		for {
			select {
			case _, ok := <-p1DrawCh:
				if ok {
					if c := p1.Draw(); c == nil {
						t.Fatal("p1 deck is empty")
					}
					t.Log("p1.Draw()")
				}
			case _, ok := <-p2DrawCh:
				if ok {
					if c := p2.Draw(); c == nil {
						t.Fatal("p1 deck is empty")
					}
					t.Log("p2.Draw()")
				}
			default:
			}
		}
	}()

	for ctx.GetState() != nil {
		if err := ctx.Exec(); err != nil {
			break
		}
	}
}
