package discard

import "github.com/hatajoe/hatastone/player/card"

type Discard struct {
	cards card.Cards
}

func NewDiscard() *Discard {
	return &Discard{}
}

func (d Discard) GetCards() card.Cards {
	return d.cards
}

func (d *Discard) Add(c card.ICard) {
	d.cards = append(card.Cards{c}, d.cards...)
}
