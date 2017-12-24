package controllers

import (
	"fmt"

	"github.com/hatajoe/hatastone/match/protocol"
)

type Draw struct{}

func (c Draw) Read() (interface{}, error) {
	return nil, nil
}

func (c Draw) Write(it interface{}) error {
	d, ok := it.(*protocol.DrawProtocol)
	if !ok {
		return fmt.Errorf("unexpected protocol specified. expected is *proxy.DrawProtocol, actual=%T", d)
	}
	if d.Card != nil {
		fmt.Printf("Player %s was drawed the Card that ID is %s\n", d.Player.GetID(), d.Card.GetID())
	}
	return nil
}
