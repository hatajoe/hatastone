package proxy

import (
	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
)

type IProxy interface {
	Listen(apps.Controller) event.IEvent
}
