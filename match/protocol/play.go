package protocol

import "github.com/hatajoe/hatastone/player/context/match"

type PlayWrite struct {
	Protocol
	Self     match.IPlayer
	Opponent match.IPlayer
}

func NewPlayWrite(self, opponent match.IPlayer, err error) *PlayWrite {
	return &PlayWrite{
		Protocol: Protocol{
			Err: err,
		},
		Self:     self,
		Opponent: opponent,
	}
}

type PlayRead struct {
	Protocol
	ID  string
	Pos int
}

func NewPlayRead(id string, pos int) *PlayRead {
	return &PlayRead{
		Protocol: Protocol{
			Err: nil,
		},
		ID:  id,
		Pos: pos,
	}
}

func (r PlayRead) GetID() string {
	return r.ID
}

func (r PlayRead) GetPos() int {
	return r.Pos
}
