package resolver

type IStrategy interface {
	Resolve(i IInfluencer, a IAccepter)
}

type IInfluencer interface {
	GetAttack() int
}

type IAccepter interface {
	AcceptAttack(damage int)
}
