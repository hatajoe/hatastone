package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/hatajoe/hatastone/match/protocol"
)

type Marigan struct{}

func (c Marigan) Read() (interface{}, error) {
	var stdin string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stdin = scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return protocol.NewMariganRead(strings.Split(stdin, " ")), nil
}

func (c Marigan) Write(it interface{}) error {
	d, ok := it.(*protocol.MariganWrite)
	if !ok {
		return fmt.Errorf("unexpected protocol specified. expected is *proxy.MariganProtocol, actual=%T", d)
	}
	if d.Err != nil {
		fmt.Println(d.Err.Error())
		return nil
	}
	fmt.Printf("Player %s marigan\n", d.Player.GetID())
	buf, err := json.Marshal(d.Player.GetHand().GetCards())
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}
