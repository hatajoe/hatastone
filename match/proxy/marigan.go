package proxy

import (
	"fmt"
	"strings"

	"github.com/hatajoe/hatastone/apps"
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

func (p *MariganProxy) Listen(r apps.Reader, w apps.Writer) event.IEvent {
	ch := make(event.Marigan)

	go func() {
		for n := range ch {
			id := []string{}
			for {
				w.Write([]byte(fmt.Sprintf("player %s marigan", p.p.GetID())))
				for _, c := range p.p.ShowHands() {
					w.Write([]byte(fmt.Sprintf("player %s card %#v", p.p.GetID(), c.GetID())))
				}
				buf, err := r.Read()
				if err != nil {
					w.Write([]byte(err.Error()))
					continue
				}
				id = strings.Split(string(buf), ",")
				c := p.p.Marigan(id)
				if len(c) <= 0 {
					w.Write([]byte(fmt.Sprintf("%s deck is empty", p.p.GetID())))
					continue
				}
				w.Write([]byte(fmt.Sprintf("%s marigan %s", p.p.GetID(), strings.Join(c.GetID(), ", "))))
				break
			}
			n.Done()
		}
	}()

	return ch
}
