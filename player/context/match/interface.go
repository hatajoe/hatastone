package match

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

type IPlayer interface {
	GetID() PlayerID
	Draw() card.ICard
	Marigan(id []string) card.Cards
	Play(id string, pos int) (card.ICard, error)
	GetHero() hero.IHero
	ShowHands() card.Cards
	FindFieldMinionByID(id string) card.IMinion
	IsHeroDead() bool
	DiscardFromHand(id string) error
	DiscardFromField(id string) error
}

type Players []IPlayer

func (ps Players) FindByID(id PlayerID) IPlayer {
	for _, p := range ps {
		if p.GetID() == id {
			return p
		}
	}
	return nil
}
