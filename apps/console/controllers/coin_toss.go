package controllers

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/protocol"
)

type CoinToss struct{}

func (c CoinToss) Read() (interface{}, error) {
	return nil, nil
}

func (c CoinToss) Write(it interface{}) error {
	d, ok := it.(*protocol.CoinTossWrite)
	if !ok {
		return fmt.Errorf("unexpected protocol specified. expected is *proxy.CoinTossProtocol, actual=%T", d)
	}
	fmt.Printf("Player %s order is %d\n", d.Player.GetID(), d.Order)
	return nil
}
