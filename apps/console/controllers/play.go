package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	cmd := strings.Split(stdin, " ")
	switch cmd[0] {
	case "play":
		if len(cmd) == 3 {
			pos, err := strconv.Atoi(cmd[2])
			if err != nil {
				return nil, err
			}
			return protocol.NewPlayRead(cmd[1], pos), nil
		}
	case "resolve":
		if len(cmd) == 4 {
			return protocol.NewResolveRead(cmd[1], cmd[2], cmd[3]), nil
		}
	default:
	}
	return nil, fmt.Errorf("command is not enough %s", stdin)
}

func (c Play) Write(it interface{}) error {
	d, ok := it.(*protocol.PlayWrite)
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
