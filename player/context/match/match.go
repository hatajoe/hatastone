package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
	"github.com/hatajoe/hatastone/player/hero"
)

type Match struct {
	ID      PlayerID
	Deck    deck.IDeck
	Hand    hand.IHand
	Field   field.IField
	Discard discard.IDiscard
}

func (m *Match) GetID() PlayerID {
	return m.ID
}

func (m *Match) Draw() card.ICard {
	return m.deckToHand()
}

func (m *Match) Marigan(id []string) card.Cards {
	ret := card.Cards{}
	for _, i := range id {
		if err := m.handToDeck(i); err != nil {
			return ret
		}
		if c := m.deckToHand(); c == nil {
			return ret
		} else {
			ret = append(ret, c)
		}
	}
	return ret
}

func (m *Match) Play(id string, pos int) (card.ICard, error) {
	return m.handToField(id, pos)
}

func (m Match) FindFieldMinionByID(id string) card.IMinion {
	return m.Field.FindByID(id)
}

func (m Match) GetHand() hand.IHand {
	return m.Hand
}

func (m *Match) DiscardFromHand(id string) error {
	return m.handToDiscard(id)
}

func (m *Match) DiscardFromField(id string) error {
	return m.fieldToDiscard(id)
}

func (m *Match) GetHero() hero.IHero {
	return m.Deck.GetHero()
}

func (m Match) IsHeroDead() bool {
	return m.Deck.GetHero().IsDead()
}

func (m *Match) handToField(id string, pos int) (card.ICard, error) {
	c := m.Hand.RemoveByID(id)
	if c == nil {
		return nil, fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	if card, ok := c.(card.IMinion); ok {
		m.Field.Add(card, pos)
		return card, nil
	}
	return nil, fmt.Errorf("unexpected card specified. id=%s", id)
}

func (m *Match) handToDeck(id string) error {
	c := m.Hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.Deck.Add(c)
	return nil
}

func (m *Match) handToDiscard(id string) error {
	c := m.Hand.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("hand.RemoveByID is failed. specified id is %s", id)
	}
	m.Discard.Add(c)
	return nil
}

func (m *Match) deckToHand() card.ICard {
	if c := m.Deck.Remove(); c != nil {
		m.Hand.Add(c)
		return c
	}
	return nil
}

func (m *Match) fieldToDiscard(id string) error {
	c := m.Field.RemoveByID(id)
	if c == nil {
		return fmt.Errorf("field.RemoveByID is failed. specified id is %s", id)
	}
	m.Discard.Add(c)
	return nil
}
