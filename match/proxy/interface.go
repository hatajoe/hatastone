package proxy

import (
	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
)

type IProxy interface {
	Listen(r apps.Reader, w apps.Writer) event.IEvent
}
