package proxy

import (
	"fmt"

	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/protocol"
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

func (p *MariganProxy) Listen(ctrl apps.Controller) event.IEvent {
	ch := make(event.Marigan)

	go func() {
		for n := range ch {
			for {
				ctrl.Write(protocol.NewMariganWrite(p.p, nil))
				it, err := ctrl.Read()
				if err != nil {
					ctrl.Write(protocol.NewMariganWrite(p.p, err))
					continue
				}
				mariganRead, ok := it.(*protocol.MariganRead)
				if !ok {
					ctrl.Write(protocol.NewMariganWrite(p.p, fmt.Errorf("unexpected type specified. expected=*protocol.MariganRead, actual=%T", mariganRead)))
					continue
				}
				c := p.p.Marigan(mariganRead.GetID())
				if len(c) <= 0 {
					ctrl.Write(protocol.NewMariganWrite(p.p, fmt.Errorf("deck is empty")))
					continue
				}
				ctrl.Write(protocol.NewMariganWrite(p.p, nil))
				break
			}
			n.Done()
		}
	}()

	return ch
}
