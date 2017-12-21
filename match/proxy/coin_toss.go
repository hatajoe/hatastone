package proxy

import (
	"log"

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

func (p *CoinTossProxy) Listen() event.IEvent {
	ch := make(event.CoinToss)

	go func() {
		for i := range ch {
			log.Printf("%s order is %d\n", p.p.GetID(), i)
		}
	}()

	return ch
}
