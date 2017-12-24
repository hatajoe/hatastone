package proxy

import (
	"fmt"

	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
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

func (p *DrawProxy) Listen(r apps.Reader, w apps.Writer) event.IEvent {
	ch := make(event.Draw)

	go func() {
		for n := range ch {
			c := p.p.Draw()
			if c == nil {
				w.Write([]byte(fmt.Sprintf("%s deck is empty", p.p.GetID())))
			}
			w.Write([]byte(fmt.Sprintf("%s draw %s", p.p.GetID(), c.GetID())))
			n.Done()
		}
	}()

	return ch
}
