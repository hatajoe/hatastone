package proxy

import "github.com/hatajoe/hatastone/match/event"

type IProxy interface {
	Listen() event.IEvent
}
