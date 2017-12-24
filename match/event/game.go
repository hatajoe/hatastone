package event

import (
	"fmt"

	player "github.com/hatajoe/hatastone/player/context/match"
)

type GameNotify struct {
	opponent player.IPlayer
	notify
}

func NewGameNotify(op player.IPlayer, done Done) *GameNotify {
	return &GameNotify{
		opponent: op,
		notify: notify{
			done: done,
		},
	}
}

func (p GameNotify) GetOpponent() player.IPlayer {
	return p.opponent
}

type Game chan *GameNotify

func GetGameEventID() EventID {
	return EventID("gameEvent")
}

func (e Game) GetID() EventID {
	return GetGameEventID()
}

func (e Game) Emit(in IEventNotify) error {
	n, ok := in.(*GameNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=GameNotify, actual=%s", n)
	}
	e <- n
	return nil
}
