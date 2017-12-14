package card

type ICard interface {
	GetID() string
}

type IEquipment interface {
	ICard
	GetAttack() int
}

type IMinion interface {
	ICard
	GetLife() int
	GetAttack() int
	AcceptAttack(damage int)
}

type ISpel interface{}
