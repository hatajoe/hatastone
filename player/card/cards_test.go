package card

import (
	"testing"
)

func TestFindByID(t *testing.T) {
	c := Cards{
		NewMurloc("murloc", 2, 2),
		NewWeapon("weapon", 2),
		NewInstant("instant", 2),
	}
	f := c.FindByID("instant")
	if f == nil {
		t.Fatal("No such card in the list")
	}
	if f.GetID() != "instant" {
		t.Fatalf("unexpected card was found. expected=%s, actual=%s", "instant", f.GetID())
	}
}

func TestDeleteByID(t *testing.T) {
	c := Cards{
		NewMurloc("murloc", 2, 2),
		NewWeapon("weapon", 2),
		NewInstant("instant", 2),
	}
	f := c.FindByID("instant")
	if f == nil {
		t.Fatal("No such card in the list")
	}
	if f.GetID() != "instant" {
		t.Fatalf("unexpected card was found. expected=%s, actual=%s", "instant", f.GetID())
	}
	c.DeleteByID("instant")
	f = c.FindByID("instant")
	if f != nil {
		t.Fatalf("unexpected card was found. id is %s", f.GetID())
	}
}
