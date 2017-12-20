package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
)

type match struct {
	id      PlayerID
	deck    deck.IDeck
	hand    hand.IHand
	field   field.IField
	discard discard.IDiscard
}

func (m *match) GetID() PlayerID {
	return m.id
}

func (m *match) Draw() card.ICard {
	return m.deckToHand()
}

func (m *match) Marigan(id []string) error {
	for _, i := range id {
		if err := m.handToDeck(i); err != nil {
			return err
		}
		if c := m.deckToHand(); c == nil {
			return fmt.Errorf("Deck is empty")
		}
	}
	return nil
}

func (m *match) Play(id string, pos int) error {
	return m.handToField(id, pos)
}

func (m *match) DiscardFromHand(id string) error {
	return m.handToDiscard(id)
}

func (m *match) DiscardFromField(id string) error {
	return m.fieldToDiscard(id)
}

func (m match) IsHeroDead() bool {
	return m.deck.GetHero().IsDead()
}

func (m *match) handToField(id string, pos int) error {
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

func (m *match) handToDeck(id string) error {
	c := m.hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.deck.Add(c)
	return nil
}

func (m *match) handToDiscard(id string) error {
	c := m.hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.discard.Add(c)
	return nil
}

func (m *match) deckToHand() card.ICard {
	if c := m.deck.Remove(); c != nil {
		m.hand.Add(c)
		return c
	}
	return nil
}

func (m *match) fieldToDiscard(id string) error {
	c := m.field.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("field.RemoveByID is failed. specified id is %s", id)
	}
	m.discard.Add(c)
	return nil
}
