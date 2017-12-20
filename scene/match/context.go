package match

import (
	"github.com/hatajoe/hatastone/event"
	"github.com/hatajoe/hatastone/match/rule"
)

type Context struct {
	rule  rule.IRule   // match rule
	first event.Events // first player event listeners
	after event.Events // after player event listeners

	state IState // game state
}

func NewContext(rule rule.IRule, first, after event.Events) *Context {
	return &Context{
		rule:  rule,
		first: first,
		after: after,
		state: &InCoinToss{},
	}
}

func (c *Context) Exec() error {
	return c.state.Exec(c)
}

func (c *Context) SetState(s IState) {
	c.state = s
}

func (c Context) GetState() IState {
	return c.state
}

func (c Context) GetSeed() int64 {
	return c.rule.GetSeed()
}

func (c Context) GetFirst() event.Events {
	return c.first
}

func (c Context) GetAfter() event.Events {
	return c.after
}

func (c *Context) SwapPlayOrder() {
	c.first, c.after = c.after, c.first
}
