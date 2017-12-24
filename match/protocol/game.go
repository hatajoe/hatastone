package protocol

import "github.com/hatajoe/hatastone/player/context/match"

type GameWrite struct {
	Protocol
	Self     match.IPlayer
	Opponent match.IPlayer
}

func NewGameWrite(self, opponent match.IPlayer, err error) *GameWrite {
	return &GameWrite{
		Protocol: Protocol{
			Err: err,
		},
		Self:     self,
		Opponent: opponent,
	}
}

type GameRead struct {
	Protocol
	id  string
	pos int
}

func NewGameRead(id string) *GameRead {
	return &GameRead{
		Protocol: Protocol{
			Err: nil,
		},
		id: id,
	}
}

func (r GameRead) GetID() string {
	return r.id
}

func (r GameRead) GetPos() int {
	return r.pos
}
