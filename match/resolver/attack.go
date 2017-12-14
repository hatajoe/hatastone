package resolver

type Attack struct{}

func (a Attack) Resolve(inf IInfluencer, acc IAccepter) {
	acc.AcceptAttack(inf.GetAttack())
}
