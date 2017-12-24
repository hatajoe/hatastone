package match

import (
	"fmt"
	"math/rand"

	"github.com/hatajoe/hatastone/match/event"
	"golang.org/x/sync/errgroup"
)

type InCoinToss struct{}

func (s InCoinToss) Exec(ctx *Context) error {
	rand.Seed(ctx.GetSeed())
	if rand.Intn(2) > 0 {
		ctx.SwapPlayOrder()
	}

	eg := errgroup.Group{}
	for i, e := range []event.Events{
		ctx.GetFirstEvent(),
		ctx.GetAfterEvent(),
	} {
		events := e
		order := i + 1
		eg.Go(func() error {
			ev := events.FindByID(event.GetCoinTossEventID())
			if ev == nil {
				return fmt.Errorf("event is nil. id is %s", event.GetCoinTossEventID())
			}
			d := make(event.Done)
			defer close(d)
			if err := ev.Emit(event.NewCoinTossNotify(order, d)); err != nil {
				return err
			}
			<-d
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	ctx.SetState(&InDraw{})
	return nil
}
