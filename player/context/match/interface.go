package match

import (
	"github.com/hatajoe/hatastone/player/card"
)

type IPlayer interface {
	GetID() PlayerID
	Draw() card.ICard
	Marigan(id []string) card.Cards
	Play(id string, pos int) error
	IsHeroDead() bool
	DiscardFromHand(id string) error
	DiscardFromField(id string) error
}
