package field

import "github.com/hatajoe/hatastone/player/card"

type IField interface {
	Add(c card.IMinion, pos int) error
	RemoveByID(id string) card.IMinion
}
