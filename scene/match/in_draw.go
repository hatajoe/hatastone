package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/event"
	"golang.org/x/sync/errgroup"
)

type InDraw struct{}

func (s InDraw) Exec(ctx *Context) error {
	eg := errgroup.Group{}
	for _, e := range []event.Events{
		ctx.GetFirst(),
		ctx.GetAfter(),
	} {
		events := e
		eg.Go(func() error {
			ev := events.FindByID(event.GetDrawEventID())
			if ev == nil {
				return fmt.Errorf("event is nil. id is %s", event.GetDrawEventID())
			}
			for i := 0; i < 3; i++ {
				if err := ev.Emit(event.DrawNotify{}); err != nil {
					return err
				}
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	ctx.SetState(&InMarigan{})
	return nil
}
