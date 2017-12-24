package hand

import "github.com/hatajoe/hatastone/player/card"

type Hand struct {
	Cards card.Cards
}

func NewHand() *Hand {
	return &Hand{}
}

func (h Hand) GetCards() card.Cards {
	return h.Cards
}

func (h *Hand) Add(c card.ICard) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) RemoveByID(id string) card.ICard {
	c := h.Cards.FindByID(id)
	h.Cards.DeleteByID(id)
	return c
}
