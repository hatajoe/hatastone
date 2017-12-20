package resolver

import (
	"testing"

	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

func TestResolveAttackMinionToHero(t *testing.T) {
	murloc := card.NewMurloc("murloc", 2, 2)
	hero := hero.NewMage(20)
	ctx := NewContext(&Attack{})
	ctx.Resolve(murloc, hero)

	if hero.GetLife() != 18 {
		t.Fatalf("unexpected result. life is expected=%d, actual=%d", 18, hero.GetLife())
	}
}

func TestResolveAttackHeroToMinion(t *testing.T) {
	murloc := card.NewMurloc("murloc", 2, 2)
	hero := hero.NewMage(20)
	ctx := NewContext(&Attack{})
	ctx.Resolve(hero, murloc)

	if murloc.GetLife() != 2 {
		t.Fatalf("unexpected result. life is expected=%d, actual=%d", 2, murloc.GetLife())
	}

	weapon := card.NewWeapon("weapon", 5)
	hero.AttachEquipment(weapon)
	ctx.Resolve(hero, murloc)

	if murloc.GetLife() != 0 {
		t.Fatalf("unexpected result. life is expected=%d, actual=%d", 0, murloc.GetLife())
	}
}

func TestResolveAttackSpelToHero(t *testing.T) {
	spel := card.NewInstant("instant", 2)
	hero := hero.NewMage(20)
	ctx := NewContext(&Attack{})
	ctx.Resolve(spel, hero)

	if hero.GetLife() != 18 {
		t.Fatalf("unexpected result. life is expected=%d, actual=%d", 18, hero.GetLife())
	}
}

func TestResolveAttackSpelToMinion(t *testing.T) {
	spel := card.NewInstant("instant", 2)
	murloc := card.NewMurloc("murloc", 2, 2)
	ctx := NewContext(&Attack{})
	ctx.Resolve(spel, murloc)

	if murloc.GetLife() != 0 {
		t.Fatalf("unexpected result. life is expected=%d, actual=%d", 2, murloc.GetLife())
	}
}
