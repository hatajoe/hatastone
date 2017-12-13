package hand

import "github.com/hatajoe/hatastone/player/card"

type Hand struct {
	cards card.Cards
}

func NewHand() *Hand {
	return &Hand{}
}

func (h Hand) GetCards() card.Cards {
	return h.cards
}

func (h *Hand) Add(c card.ICard) {
	h.cards = append(h.cards, c)
}

func (h *Hand) RemoveByID(id string) card.ICard {
	c := h.cards.FindByID(id)
	h.cards.DeleteByID(id)
	return c
}
