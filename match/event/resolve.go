package event

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/resolver"
)

type ResolveNotify struct {
	strategy   resolver.IStrategy
	influencer resolver.IInfluencer
	acceptor   resolver.IAcceptor
	notify
}

func NewResolveNotify(s resolver.IStrategy, i resolver.IInfluencer, a resolver.IAcceptor, done Done) *ResolveNotify {
	return &ResolveNotify{
		strategy:   s,
		influencer: i,
		acceptor:   a,
		notify: notify{
			done: done,
		},
	}
}

func (n ResolveNotify) GetStrategy() resolver.IStrategy {
	return n.strategy
}

func (n ResolveNotify) GetInfluencer() resolver.IInfluencer {
	return n.influencer
}

func (n ResolveNotify) GetAcceptor() resolver.IAcceptor {
	return n.acceptor
}

type Resolve chan *ResolveNotify

func GetResolveEventID() EventID {
	return EventID("resolveEvent")
}

func (e Resolve) GetID() EventID {
	return GetResolveEventID()
}

func (e Resolve) Emit(in IEventNotify) error {
	n, ok := in.(*ResolveNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=ResolveNotify, actual=%s", n)
	}
	e <- n
	return nil
}
