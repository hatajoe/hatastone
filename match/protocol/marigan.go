package protocol

import (
	"github.com/hatajoe/hatastone/player/context/match"
)

type MariganWrite struct {
	Protocol
	Player match.IPlayer
}

func NewMariganWrite(player match.IPlayer, err error) *MariganWrite {
	return &MariganWrite{
		Protocol: Protocol{
			Err: err,
		},
		Player: player,
	}
}

type MariganRead struct {
	id []string
}

func NewMariganRead(id []string) *MariganRead {
	return &MariganRead{
		id: id,
	}
}

func (r MariganRead) GetID() []string {
	return r.id
}
