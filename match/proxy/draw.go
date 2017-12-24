package proxy

import (
	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/protocol"
	"github.com/hatajoe/hatastone/player/context/match"
)

type DrawProxy struct {
	proxy
}

func NewDrawProxy(p match.IPlayer) *DrawProxy {
	return &DrawProxy{
		proxy: proxy{
			p: p,
		},
	}
}

func (p *DrawProxy) Listen(ctrl apps.Controller) event.IEvent {
	ch := make(event.Draw)

	go func() {
		for n := range ch {
			c := p.p.Draw()
			ctrl.Write(protocol.NewDrawProtocol(p.p, c))
			n.Done()
		}
	}()

	return ch
}
