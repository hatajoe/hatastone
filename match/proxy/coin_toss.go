package proxy

import (
	"fmt"

	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
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

func (p *CoinTossProxy) Listen(r apps.Reader, w apps.Writer) event.IEvent {
	ch := make(event.CoinToss)

	go func() {
		for n := range ch {
			w.Write([]byte(fmt.Sprintf("%s order is %d\n", p.p.GetID(), n.GetOrder())))
			n.Done()
		}
	}()

	return ch
}
