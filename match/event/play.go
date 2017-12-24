package event

import (
	"fmt"

	player "github.com/hatajoe/hatastone/player/context/match"
)

type PlayNotify struct {
	opponent player.IPlayer
	events   Events
	notify
}

func NewPlayNotify(op player.IPlayer, events Events, done Done) *PlayNotify {
	return &PlayNotify{
		opponent: op,
		events:   events,
		notify: notify{
			done: done,
		},
	}
}

func (p PlayNotify) GetOpponent() player.IPlayer {
	return p.opponent
}

func (p PlayNotify) GetEvents() Events {
	return p.events
}

type Play chan *PlayNotify

func GetPlayEventID() EventID {
	return EventID("playEvent")
}

func (e Play) GetID() EventID {
	return GetPlayEventID()
}

func (e Play) Emit(in IEventNotify) error {
	n, ok := in.(*PlayNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=PlayNotify, actual=%s", n)
	}
	e <- n
	return nil
}
