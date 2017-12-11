package deck

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

type Deck struct {
	hero  hero.IHero
	cards []card.ICard
}

func NewDeck(h hero.IHero, c []card.ICard) *Deck {
	return &Deck{
		hero:  h,
		cards: c,
	}
}

func (d *Deck) Push(c card.ICard) {
	d.cards = append(d.cards, c)
}

func (d *Deck) Pop() card.ICard {
	if len(d.cards) <= 0 {
		return nil
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}
