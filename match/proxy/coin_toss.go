package proxy

import (
	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/protocol"
	"github.com/hatajoe/hatastone/player/context/match"
)

type CoinTossProxy struct {
	proxy
}

func NewCoinTossProxy(p match.IPlayer) *CoinTossProxy {
	return &CoinTossProxy{
		proxy: proxy{
			p: p,
		},
	}
}

func (p *CoinTossProxy) Listen(ctrl apps.Controller) event.IEvent {
	ch := make(event.CoinToss)

	go func() {
		for n := range ch {
			ctrl.Write(protocol.NewCoinTossProtocol(p.p, n.GetOrder()))
			n.Done()
		}
	}()

	return ch
}
