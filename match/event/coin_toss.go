package event

import "fmt"

type CoinTossNotify int

type CoinToss chan CoinTossNotify

func GetCoinTossEventID() EventID {
	return EventID("coinTossEvent")
}

func (e CoinToss) GetID() EventID {
	return GetCoinTossEventID()
}

func (e CoinToss) Emit(in IEventNotify) error {
	n, ok := in.(CoinTossNotify)
	if !ok {
		return fmt.Errorf("unexpected event notify specified. expected=event.CoinTossNotify, actual=%T", n)
	}
	e <- n
	return nil
}
