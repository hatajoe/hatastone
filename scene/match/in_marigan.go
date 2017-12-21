package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/event"
	"golang.org/x/sync/errgroup"
)

type InMarigan struct{}

func (s InMarigan) Exec(ctx *Context) error {
	eg := errgroup.Group{}
	for _, e := range []event.Events{
		ctx.GetFirst(),
		ctx.GetAfter(),
	} {
		events := e
		eg.Go(func() error {
			ev := events.FindByID(event.GetMariganEventID())
			if ev == nil {
				return fmt.Errorf("event is nil. id is %s", event.GetMariganEventID())
			}
			if err := ev.Emit(event.MariganNotify{}); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	ctx.SetState(&InPlay{})
	return nil
}
