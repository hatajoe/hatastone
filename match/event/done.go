package event

import (
	"fmt"
)

type DoneNotify struct{}

type Done chan *DoneNotify

func GetDoneEventID() EventID {
	return EventID("doneEvent")
}

func (e Done) GetID() EventID {
	return GetDoneEventID()
}

func (e Done) Emit(in IEventNotify) error {
	n, ok := in.(*DoneNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=DoneNotify, actual=%s", n)
	}
	e <- n
	return nil
}
