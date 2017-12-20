package rule

type Rule struct {
	seed        int64
	initialHand int
}

func NewRule(seed int64, initialHand int) *Rule {
	return &Rule{
		seed:        seed,
		initialHand: initialHand,
	}
}

func (r Rule) GetSeed() int64 {
	return r.seed
}

func (r Rule) GetInitialHand() int {
	return r.initialHand
}
