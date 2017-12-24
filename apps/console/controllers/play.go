package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hatajoe/hatastone/match/protocol"
)

type Play struct{}

func (c Play) Read() (interface{}, error) {
	var stdin string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stdin = scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return protocol.NewGameRead(stdin), nil
}

func (c Play) Write(it interface{}) error {
	d, ok := it.(*protocol.GameWrite)
	if !ok {
		return fmt.Errorf("unexpected protocol specified. expected is *proxy.PlayProtocol, actual=%T", d)
	}
	if d.Err != nil {
		fmt.Println(d.Err.Error())
		return nil
	}
	buf, err := json.Marshal(d.Self.GetHero())
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	buf, err = json.Marshal(d.Self.GetHand().GetCards())
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}
