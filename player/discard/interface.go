package discard

import "github.com/hatajoe/hatastone/player/card"

type IDiscard interface {
	Add(c card.ICard)
}
