package deck

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/card/equipment"
	"github.com/hatajoe/hatastone/player/card/minion"
	"github.com/hatajoe/hatastone/player/card/spel"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestPop(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{
			minion.NewMurloc(),
			equipment.NewWeapon(),
			spel.NewInstant(),
		},
	)
	if actual, ok := d.Pop().(*minion.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &minion.Murloc{}, actual)
	}
	if actual, ok := d.Pop().(*equipment.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &equipment.Weapon{}, actual)
	}
}

func TestPush(t *testing.T) {
	d := NewDeck(
		hero.NewMage(),
		[]card.ICard{},
	)
	d.Push(minion.NewMurloc())
	d.Push(equipment.NewWeapon())
	if actual, ok := d.Pop().(*minion.Murloc); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &minion.Murloc{}, actual)
	}
	if actual, ok := d.Pop().(*equipment.Weapon); !ok {
		t.Fatalf("unexpected card is popped. expected=%T, actual=%T", &equipment.Weapon{}, actual)
	}
}
