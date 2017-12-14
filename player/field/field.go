package field

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
)

type Field struct {
	cards card.Cards
}

func NewField() *Field {
	return &Field{}
}

func (f Field) GetCards() card.Cards {
	return f.cards
}

func (f *Field) Add(c card.ICard, pos int) error {
	if pos < 0 || len(f.cards) < pos {
		return fmt.Errorf("out of range error. expected <= %d, actual=%d", len(f.cards), pos)
	}
	f.cards = append(f.cards[:pos], append(card.Cards{c}, f.cards[pos:]...)...)
	return nil
}

func (f *Field) RemoveByID(id string) card.ICard {
	c := f.cards.FindByID(id)
	f.cards.DeleteByID(id)
	return c
}
