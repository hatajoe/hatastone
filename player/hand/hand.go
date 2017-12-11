package hand

import "github.com/hatajoe/hatastone/player/card"

type Hand struct {
	cards []card.ICard
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) Push(c card.ICard) {
	h.cards = append(h.cards, c)
}

func (h *Hand) PopByIndex(idx int) card.ICard {
	l := len(h.cards)
	if l <= 0 || idx < 0 || idx > l {
		return nil
	}
	c := h.cards[idx]
	h.cards = append(h.cards[0:idx], h.cards[idx:])
	return c
}
