package protocol

import "github.com/hatajoe/hatastone/player/context/match"

type ResolveWrite struct {
	Protocol
	Self     match.IPlayer
	Opponent match.IPlayer
}

func NewResolveWrite(self, opponent match.IPlayer, err error) *ResolveWrite {
	return &ResolveWrite{
		Protocol: Protocol{
			Err: err,
		},
		Self:     self,
		Opponent: opponent,
	}
}

type ResolveRead struct {
	Protocol
	Cmd          string
	InfluencerID string
	AcceptorID   string
}

func NewResolveRead(cmd string, influencerID string, acceptorID string) *ResolveRead {
	return &ResolveRead{
		Protocol: Protocol{
			Err: nil,
		},
		Cmd:          cmd,
		InfluencerID: influencerID,
		AcceptorID:   acceptorID,
	}
}

func (r ResolveRead) GetCmd() string {
	return r.Cmd
}

func (r ResolveRead) GetInfluencerID() string {
	return r.InfluencerID
}

func (r ResolveRead) GetAcceptorID() string {
	return r.AcceptorID
}
