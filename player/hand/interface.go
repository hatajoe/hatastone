package hand

import "github.com/hatajoe/hatastone/player/card"

type IHand interface {
	Push(c card.ICard)
	PopByID(idx string) card.ICard
}
