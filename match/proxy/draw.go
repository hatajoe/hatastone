package proxy

import (
	"fmt"
	"log"

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

func (p *DrawProxy) Listen() event.IEvent {
	ch := make(event.Draw)

	go func() {
		for range ch {
			c := p.p.Draw()
			if c == nil {
				p.err <- fmt.Errorf("%s deck is empty", p.p.GetID())
			}
			log.Printf("%s draw %s\n", p.p.GetID(), c.GetID())
		}
	}()

	return ch
}
