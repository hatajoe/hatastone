package discard

import "github.com/hatajoe/hatastone/player/card"

type Discard struct {
	Cards card.Cards
}

func NewDiscard() *Discard {
	return &Discard{}
}

func (d Discard) GetCards() card.Cards {
	return d.Cards
}

func (d *Discard) Add(c card.ICard) {
	d.Cards = append(card.Cards{c}, d.Cards...)
}
