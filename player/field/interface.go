package field

import "github.com/hatajoe/hatastone/player/card"

type IField interface {
	Add(c card.ICard, pos int) error
	RemoveByID(id string) card.ICard
}
