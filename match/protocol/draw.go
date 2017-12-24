package protocol

import (
	"github.com/hatajoe/hatastone/player/card"
	"github.com/hatajoe/hatastone/player/context/match"
)

type DrawProtocol struct {
	Protocol
	Player match.IPlayer
	Card   card.ICard
}

func NewDrawProtocol(p match.IPlayer, c card.ICard) *DrawProtocol {
	return &DrawProtocol{
		Protocol: Protocol{
			Err: nil,
		},
		Player: p,
		Card:   c,
	}
}
