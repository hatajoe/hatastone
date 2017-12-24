package deck

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

type Deck struct {
	Hero  hero.IHero
	Cards card.Cards
}

func NewDeck(h hero.IHero, c []card.ICard) *Deck {
	return &Deck{
		Hero:  h,
		Cards: c,
	}
}

func (d Deck) GetHero() hero.IHero {
	return d.Hero
}

func (d *Deck) Add(c card.ICard) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Remove() card.ICard {
	if len(d.Cards) <= 0 {
		return nil
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}
