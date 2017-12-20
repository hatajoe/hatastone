package hero

type Mage struct {
	Hero
}

func NewMage(life int) *Mage {
	return &Mage{
		Hero: Hero{
			life: life,
		},
	}
}
