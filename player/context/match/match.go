package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
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

func (m *Match) Marigan(id []string) error {
	for _, i := range id {
		if err := m.handToDeck(i); err != nil {
			return err
		}
		if !m.deckToHand() {
			return fmt.Errorf("Deck is empty")
		}
	}
	return nil
}

func (m *Match) Play(id string, pos int) error {
	return m.handToField(id, pos)
}

func (m *Match) DiscardFromHand(id string) error {
	return m.handToDiscard(id)
}

func (m *Match) DiscardFromField(id string) error {
	return m.fieldToDiscard(id)
}

func (m *Match) handToField(id string, pos int) error {
	c := m.hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	if card, ok := c.(card.IMinion); ok {
		m.field.Add(card, pos)
		return nil
	}
	return fmt.Errorf("unexpected card specified. id=%s", id)
}

func (m *Match) handToDeck(id string) error {
	c := m.hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.deck.Add(c)
	return nil
}

func (m *Match) handToDiscard(id string) error {
	c := m.hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.discard.Add(c)
	return nil
}

func (m *Match) deckToHand() bool {
	if c := m.deck.Remove(); c != nil {
		m.hand.Add(c)
	} else {
		return false
	}
	return true
}

func (m *Match) fieldToDiscard(id string) error {
	c := m.field.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("field.RemoveByID is failed. specified id is %s", id)
	}
	m.discard.Add(c)
	return nil
}
