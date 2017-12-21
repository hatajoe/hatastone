package proxy

import (
	"fmt"
	"log"
	"strings"

	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/player/context/match"
)

type MariganProxy struct {
	proxy
}

func NewMariganProxy(p match.IPlayer) *MariganProxy {
	return &MariganProxy{
		proxy: proxy{
			p: p,
		},
	}
}

func (p *MariganProxy) Listen() event.IEvent {
	ch := make(event.Marigan)

	go func() {
		for range ch {
			c := p.p.Marigan([]string{"p1m1", "p1i1"})
			if len(c) <= 0 {
				p.err <- fmt.Errorf("%s deck is empty", p.p.GetID())
			}
			log.Printf("%s marigan %s", p.p.GetID(), strings.Join(c.GetID(), ", "))
		}
	}()

	return ch
}
