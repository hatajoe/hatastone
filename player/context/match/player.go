package match

import (
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
)

type Player struct {
	match
}

func NewPlayer(id PlayerID, dc deck.IDeck, hn hand.IHand, fl field.IField, ds discard.IDiscard) *Player {
	return &Player{
		match: match{
			id:      id,
			deck:    dc,
			hand:    hn,
			field:   fl,
			discard: ds,
		},
	}
}
