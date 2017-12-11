package deck

import "github.com/hatajoe/hatastone/player/card"

type IDeck interface {
	Pop() card.ICard
	Push(c card.ICard)
}
