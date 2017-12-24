package proxy

import (
	"github.com/hatajoe/hatastone/apps"
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/resolver"
)

type ResolveProxy struct {
	err chan error
}

func NewResolveProxy(err chan error) *ResolveProxy {
	return &ResolveProxy{
		err: err,
	}
}

func (p *ResolveProxy) Listen(r apps.Reader, w apps.Writer) event.IEvent {
	ch := make(event.Resolve)

	go func() {
		for n := range ch {
			ctx := resolver.NewContext(n.GetStrategy())
			ctx.Resolve(n.GetInfluencer(), n.GetAcceptor())
			n.Done()
		}
	}()

	return ch
}
