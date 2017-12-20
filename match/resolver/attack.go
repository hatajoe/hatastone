package resolver

type Attack struct{}

func (a Attack) Resolve(inf IInfluencer, acc IAcceptor) {
	acc.AcceptAttack(inf.GetAttack())
}
