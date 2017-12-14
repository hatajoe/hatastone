package field

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
)

func TestAdd(t *testing.T) {
	expected := card.Cards{
		card.NewMurloc("murloc1", 2, 2),
		card.NewMurloc("murloc2", 2, 2),
		card.NewMurloc("murloc3", 2, 2),
	}

	f := NewField()
	f.Add(card.NewMurloc("murloc1", 2, 2), 0)
	f.Add(card.NewMurloc("murloc3", 2, 2), 1)
	f.Add(card.NewMurloc("murloc2", 2, 2), 1)

	for i, card := range f.GetMinions() {
		if expected[i].GetID() != card.GetID() {
			t.Errorf("field.Add is unexpected behavior. expected=%s, actual=%s", expected[i].GetID(), card.GetID())
		}
	}
}

func TestRemoveByID(t *testing.T) {
	f := NewField()
	if err := f.Add(card.NewMurloc("murloc1", 2, 2), 1); err == nil {
		t.Fatalf("error shuould be occured")
	}
	if err := f.Add(card.NewMurloc("murloc1", 2, 2), 0); err != nil {
		t.Fatalf("unexpected error occured. %s", err)
	}
	if actual, ok := f.RemoveByID("murloc1").(*card.Murloc); !ok {
		t.Fatalf("unexpected card is removed. expected=%T, actual=%T", &card.Murloc{}, actual)
	}
	if c := f.RemoveByID("murloc1"); c != nil {
		t.Fatalf("unexpected value removed. expected=nil, actual=%T", c)
	}
	if c := f.RemoveByID("pirates"); c != nil {
		t.Fatalf("unexpected value removed. expected=nil, actual=%T", c)
	}
}
