package protocol

import "github.com/hatajoe/hatastone/player/context/match"

type CoinTossProtocol struct {
	Protocol
	Player match.IPlayer
	Order  int
}

func NewCoinTossProtocol(p match.IPlayer, order int) *CoinTossProtocol {
	return &CoinTossProtocol{
		Protocol: Protocol{
			Err: nil,
		},
		Player: p,
		Order:  order,
	}
}
