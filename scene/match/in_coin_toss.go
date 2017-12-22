package match

import (
	"fmt"
	"math/rand"

	"github.com/hatajoe/hatastone/match/event"
)

type InCoinToss struct{}

func (s InCoinToss) Exec(ctx *Context) error {
	rand.Seed(ctx.GetSeed())
	if rand.Intn(2) > 0 {
		ctx.SwapPlayOrder()
	}

	for i, e := range []event.Events{
		ctx.GetFirst(),
		ctx.GetAfter(),
	} {
		ev := e.FindByID(event.GetCoinTossEventID())
		if ev == nil {
			return fmt.Errorf("event is nil. id is %s", event.GetCoinTossEventID())
		}
		d := make(event.Done)
		if err := ev.Emit(event.NewCoinTossNotify(i, d)); err != nil {
			return err
		}
		<-d
	}

	ctx.SetState(&InDraw{})
	return nil
}
