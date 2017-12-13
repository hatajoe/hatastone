package deck

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestRemove(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{
			&card.Murloc{},
			&card.Weapon{},
			&card.Instant{},
		},
	)
	if actual, ok := d.Remove().(*card.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Murloc{}, actual)
	}
	if actual, ok := d.Remove().(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}

func TestAdd(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{},
	)
	d.Add(&card.Murloc{})
	d.Add(&card.Weapon{})
	if actual, ok := d.Remove().(*card.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Murloc{}, actual)
	}
	if actual, ok := d.Remove().(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}
