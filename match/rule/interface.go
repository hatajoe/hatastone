package rule

type IRule interface {
	GetSeed() int64
	GetInitialHand() int
}
