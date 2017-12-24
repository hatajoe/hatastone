package proxy

import (
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
			ctrl.Write(protocol.NewPlayWrite(p.p, n.GetOpponent(), nil))
			it, err := ctrl.Read()
			if err != nil {
				ctrl.Write(protocol.NewPlayWrite(nil, nil, err))
				n.Done()
				continue
			}

			switch reader := it.(type) {
			case *protocol.PlayRead:
				_, err = p.p.Play(reader.GetID(), reader.GetPos())
				if err != nil {
					ctrl.Write(protocol.NewPlayWrite(nil, nil, err))
					n.Done()
					continue
				}
			case *protocol.ResolveRead:
				resolver.NewContext(&resolver.Attack{}).Resolve(p.p.FindFieldMinionByID(reader.GetInfluencerID()), n.GetOpponent().GetHero())
			default:
				ctrl.Write(protocol.NewPlayWrite(nil, nil, err))
				n.Done()
				continue
			}
			n.Done()
			ctrl.Write(protocol.NewPlayWrite(p.p, n.GetOpponent(), nil))
		}
	}()

	return ch
}
