package hand

import "github.com/hatajoe/hatastone/player/card"

type IHand interface {
	Add(c card.ICard)
	RemoveByID(idx string) card.ICard
}
