package resolver

type IStrategy interface {
	Resolve(i IInfluencer, a IAcceptor)
}

type IInfluencer interface {
	GetAttack() int
}

type IAcceptor interface {
	AcceptAttack(damage int)
}
