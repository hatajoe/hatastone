package match

import (
	"fmt"
	"log"

	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/proxy"
)

type InPlay struct{}

func (s InPlay) Exec(ctx *Context) error {
	for i := 0; i < 2; i++ {
		e := ctx.GetCurrentEvents()

		if err := draw(e); err != nil {
			return err
		}

		ev := e.FindByID(event.GetPlayEventID())
		if ev == nil {
			return fmt.Errorf("event is nil. id is %s", event.GetPlayEventID())
		}
		d := make(event.Done)
		if err := ev.Emit(event.NewPlayNotify(ctx.GetOpponent(), event.Events{
			resolve(),
		}, d)); err != nil {
			close(d)
			return err
		}
		<-d
		close(d)

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

func resolve() event.IEvent {
	err := make(chan error)
	go func() {
		for e := range err {
			log.Println(e)
		}
	}()
	resolver := proxy.NewResolveProxy(err)

	return resolver.Listen(nil, nil)
}
