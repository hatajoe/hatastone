package event

import "fmt"

type MariganNotify struct {
	notify
}

type Marigan chan MariganNotify

func GetMariganEventID() EventID {
	return EventID("mariganEvent")
}

func (e Marigan) GetID() EventID {
	return GetMariganEventID()
}

func (e Marigan) Emit(in IEventNotify) error {
	n, ok := in.(MariganNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=MariganNotify, actual=%s", n)
	}
	e <- n
	return nil
}
