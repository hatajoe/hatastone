package deck

import "github.com/hatajoe/hatastone/player/card"

type IDeck interface {
	Add(c card.ICard)
	Remove() card.ICard
}
