package deck

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestPop(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{
			&card.Murloc{},
			&card.Weapon{},
			&card.Instant{},
		},
	)
	if actual, ok := d.Pop().(*card.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Murloc{}, actual)
	}
	if actual, ok := d.Pop().(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}

func TestPush(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{},
	)
	d.Push(&card.Murloc{})
	d.Push(&card.Weapon{})
	if actual, ok := d.Pop().(*card.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Murloc{}, actual)
	}
	if actual, ok := d.Pop().(*card.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &card.Weapon{}, actual)
	}
}
