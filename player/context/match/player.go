package match

import (
	"github.com/hatajoe/hatastone/player/deck"
	"github.com/hatajoe/hatastone/player/discard"
	"github.com/hatajoe/hatastone/player/field"
	"github.com/hatajoe/hatastone/player/hand"
)

type Player struct {
	Match
}

func NewPlayer(id PlayerID, dc deck.IDeck, hn hand.IHand, fl field.IField, ds discard.IDiscard) *Player {
	return &Player{
		Match: Match{
			ID:      id,
			Deck:    dc,
			Hand:    hn,
			Field:   fl,
			Discard: ds,
		},
	}
}
