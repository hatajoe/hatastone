package proxy

import (
	"fmt"

	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/protocol"
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

func (p *PlayProxy) Listen(ctrl apps.Controller) event.IEvent {
	ch := make(event.Game)

	go func() {
		for n := range ch {
			ctrl.Write(protocol.NewGameWrite(p.p, n.GetOpponent(), nil))
			it, err := ctrl.Read()
			if err != nil {
				ctrl.Write(protocol.NewGameWrite(nil, nil, err))
				n.Done()
				continue
			}
			reader, ok := it.(*protocol.GameRead)
			if !ok {
				ctrl.Write(protocol.NewGameWrite(nil, nil, fmt.Errorf("unexpected type specified. expected=*protocol.PlayRead, actual=%T", reader)))
				n.Done()
				continue
			}

			_, err = p.p.Play(reader.GetID(), reader.GetPos())
			if err != nil {
				ctrl.Write(protocol.NewGameWrite(nil, nil, err))
				n.Done()
				continue
			}

			events := n.GetEvents()
			ev := events.FindByID(event.GetResolveEventID())
			if ev == nil {
				ctrl.Write(protocol.NewGameWrite(nil, nil, fmt.Errorf("event is nil. id is %s", event.GetResolveEventID())))
				n.Done()
				continue
			}
			d := make(event.Done)
			if err := ev.Emit(event.NewResolveNotify(&resolver.Attack{}, p.p.FindFieldMinionByID(reader.GetID()), n.GetOpponent().GetHero(), d)); err != nil {
				ctrl.Write(protocol.NewGameWrite(nil, nil, err))
			}
			<-d
			close(d)
			n.Done()
			ctrl.Write(protocol.NewGameWrite(p.p, n.GetOpponent(), nil))
		}
	}()

	return ch
}
