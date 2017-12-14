package field

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
)

type Field struct {
	cards card.Minions
}

func NewField() *Field {
	return &Field{}
}

func (f Field) GetMinions() card.Minions {
	return f.cards
}

func (f *Field) Add(c card.IMinion, pos int) error {
	if pos < 0 || len(f.cards) < pos {
		return fmt.Errorf("out of range error. expected <= %d, actual=%d", len(f.cards), pos)
	}
	f.cards = append(f.cards[:pos], append(card.Minions{c}, f.cards[pos:]...)...)
	return nil
}

func (f *Field) RemoveByID(id string) card.IMinion {
	c := f.cards.FindByID(id)
	f.cards.DeleteByID(id)
	return c
}
