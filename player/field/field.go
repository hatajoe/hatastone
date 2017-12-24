package field

import (
	"fmt"

	"github.com/hatajoe/hatastone/player/card"
)

type Field struct {
	Cards card.Minions
}

func NewField() *Field {
	return &Field{}
}

func (f Field) GetMinions() card.Minions {
	return f.Cards
}

func (f *Field) Add(c card.IMinion, pos int) error {
	if pos < 0 || len(f.Cards) < pos {
		return fmt.Errorf("out of range error. expected <= %d, actual=%d", len(f.Cards), pos)
	}
	f.Cards = append(f.Cards[:pos], append(card.Minions{c}, f.Cards[pos:]...)...)
	return nil
}

func (f *Field) RemoveByID(id string) card.IMinion {
	c := f.Cards.FindByID(id)
	f.Cards.DeleteByID(id)
	return c
}

func (f *Field) FindByID(id string) card.IMinion {
	for _, m := range f.Cards {
		if m.GetID() == id {
			return m
		}
	}
	return nil
}
