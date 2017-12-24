package resolver

type IStrategy interface {
	Resolve(i IInfluencer, a IAcceptor)
}

type IInfluencer interface {
	GetID() string
	GetAttack() int
}

type IAcceptor interface {
	GetID() string
	AcceptAttack(damage int)
}
