package proxy

import (
	"log"

	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/player/context/match"
)

type CoinTossProxy struct {
	proxy
}

func NewCoinTossProxy(p match.IPlayer, err chan error) *CoinTossProxy {
	return &CoinTossProxy{
		proxy: proxy{
			p:   p,
			err: err,
		},
	}
}

func (p *CoinTossProxy) Listen() event.IEvent {
	ch := make(event.CoinToss)

	go func() {
		for n := range ch {
			log.Printf("%s order is %d\n", p.p.GetID(), n.GetOrder())
			n.Done()
		}
	}()

	return ch
}
