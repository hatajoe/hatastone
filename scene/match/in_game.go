package match

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/event"
)

type InGame struct{}

func (s InGame) Exec(ctx *Context) error {
	for !ctx.IsHeroDead() {
		e := ctx.GetCurrentEvents()

		if err := draw(e); err != nil {
			return err
		}

		for {
			ev := e.FindByID(event.GetGameEventID())
			if ev == nil {
				return fmt.Errorf("event is nil. id is %s", event.GetGameEventID())
			}
			d := make(event.Done)
			if err := ev.Emit(event.NewGameNotify(ctx.GetOpponent(), d)); err != nil {
				close(d)
				continue
			}
			<-d
			close(d)
			break
		}

		ctx.SwitchCurrentPlayer()
	}

	ctx.SetState(&InResult{})
	return nil
}

func draw(events event.Events) error {
	ev := events.FindByID(event.GetDrawEventID())
	if ev == nil {
		return fmt.Errorf("event is nil. id is %s", event.GetDrawEventID())
	}
	d := make(event.Done)
	if err := ev.Emit(event.NewDrawNotify(d)); err != nil {
		return err
	}
	<-d
	close(d)
	return nil
}
