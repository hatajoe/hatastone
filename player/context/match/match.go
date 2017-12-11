package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
)

type Match struct {
	deck    deck.IDeck
	hand    hand.IHand
	field   field.IField
	discard discard.IDiscard
}

func NewMatch(dc deck.IDeck, hn hand.IHand, fl field.IField, ds discard.IDiscard) *Match {
	return &Match{
		deck:    dc,
		hand:    hn,
		field:   fl,
		discard: ds,
	}
}

func (m *Match) Draw() bool {
	return m.deckToHand()
}

func (m *Match) Marigan(idx []int) error {
	for _, i := range idx {
		if err := m.handToDeck(i); err != nil {
			return err
		}
		if !m.deckToHand() {
			return fmt.Errorf("Deck is empty")
		}
	}
	return nil
}

func (m *Match) handToDeck(idx int) error {
	c := m.hand.PopByIndex(idx)
	if c == nil {
		return fmt.Errorf("PopByIndex is failed. specified idx is %d", idx)
	}
	m.deck.Push(c)
	return nil
}

func (m *Match) deckToHand() bool {
	if c := m.deck.Pop(); c != nil {
		m.hand.Push(c)
	} else {
		return false
	}
	return true
}
