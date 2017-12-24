package protocol

import "github.com/hatajoe/hatastone/player/context/match"

type CoinTossWrite struct {
	Protocol
	Player match.IPlayer
	Order  int
}

func NewCoinTossWrite(p match.IPlayer, order int) *CoinTossWrite {
	return &CoinTossWrite{
		Protocol: Protocol{
			Err: nil,
		},
		Player: p,
		Order:  order,
	}
}
