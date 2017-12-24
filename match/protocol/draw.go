package protocol

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/context/match"
)

type DrawWrite struct {
	Protocol
	Player match.IPlayer
	Card   card.ICard
}

func NewDrawWrite(p match.IPlayer, c card.ICard) *DrawWrite {
	return &DrawWrite{
		Protocol: Protocol{
			Err: nil,
		},
		Player: p,
		Card:   c,
	}
}
