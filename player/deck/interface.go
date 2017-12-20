package deck

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/hero"
)

type IDeck interface {
	Add(c card.ICard)
	Remove() card.ICard
	GetHero() hero.IHero
}
