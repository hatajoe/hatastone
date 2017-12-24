package hero

import "github.com/hatajoe/hatastone/player/card"

type IHero interface {
	GetID() string
	GetLife() int
	IsDead() bool
	GetEquipment() card.IEquipment
	AttachEquipment(e card.IEquipment)
	DetachEquipment()
	GetAttack() int
	AcceptAttack(damage int)
}
