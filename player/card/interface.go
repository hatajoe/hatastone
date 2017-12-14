package card

type ICard interface {
	GetID() string
}

type IEquipment interface{}

type IMinion interface {
	ICard
}

type ISpel interface{}
