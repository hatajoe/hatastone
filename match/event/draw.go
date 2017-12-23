package event

import "fmt"

type DrawNotify struct {
	notify
}

func NewDrawNotify(done Done) *DrawNotify {
	return &DrawNotify{
		notify: notify{
			done: done,
		},
	}
}

type Draw chan *DrawNotify

func GetDrawEventID() EventID {
	return EventID("drawEvent")
}

func (e Draw) GetID() EventID {
	return GetDrawEventID()
}

func (e Draw) Emit(in IEventNotify) error {
	n, ok := in.(*DrawNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=DrawNotify, actual=%s", n)
	}
	e <- n
	return nil
}
