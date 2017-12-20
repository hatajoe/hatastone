package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/event"
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
			ev.Open()
			defer ev.Close()

			for i := 0; i < 3; i++ {
				ev.Emit()
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
