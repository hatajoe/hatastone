package proxy

import (
	"fmt"
	"strings"

	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/resolver"
	"github.com/hatajoe/hatastone/player/context/match"
)

type PlayProxy struct {
	proxy
}

func NewPlayProxy(p match.IPlayer) *PlayProxy {
	return &PlayProxy{
		proxy: proxy{
			p: p,
		},
	}
}

func (p *PlayProxy) Listen(r apps.Reader, w apps.Writer) event.IEvent {
	ch := make(event.Play)

	go func() {
		for n := range ch {
			w.Write([]byte(fmt.Sprintf("player %s play", p.p.GetID())))
			for _, c := range p.p.ShowHands() {
				w.Write([]byte(fmt.Sprintf("player %s hands %#v", p.p.GetID(), c.GetID())))
			}
			buf, err := r.Read()
			if err != nil {
				w.Write([]byte(err.Error()))
				continue
			}
			id := strings.Split(string(buf), ",")

			card, err := p.p.Play(id[0], 0)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(fmt.Sprintf("%s play %s\n", p.p.GetID(), card.GetID())))

			events := n.GetEvents()
			ev := events.FindByID(event.GetResolveEventID())
			if ev == nil {
				w.Write([]byte(fmt.Sprintf("event is nil. id is %s", event.GetResolveEventID())))
			}
			d := make(event.Done)
			if err := ev.Emit(event.NewResolveNotify(&resolver.Attack{}, p.p.FindFieldMinionByID(id[1]), n.GetOpponent().GetHero(), d)); err != nil {
				w.Write([]byte(err.Error()))
			}
			<-d
			close(d)
			n.Done()
			w.Write([]byte(fmt.Sprintf("player %s's hero life is %d", p.p.GetID(), p.p.GetHero().GetLife())))
			w.Write([]byte(fmt.Sprintf("player %s's hero life is %d", n.GetOpponent().GetID(), n.GetOpponent().GetHero().GetLife())))
		}
	}()

	return ch
}
