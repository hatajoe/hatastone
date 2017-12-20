package match

import "math/rand"

type InCoinToss struct{}

func (s InCoinToss) Exec(ctx *Context) error {
	rand.Seed(ctx.GetSeed())
	if rand.Intn(2) > 0 {
		ctx.SwapPlayOrder()
	}
	ctx.SetState(&InDraw{})
	return nil
}
