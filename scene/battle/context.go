package battle

import (
	"github.com/hatajoe/hatastone/battle"
)

// Context saves the state of battle
type Context struct {
	btl   *battle.Battle
	state IState
}

// NewContext returns Context instance
func NewContext(b *battle.Battle) *Context {
	return &Context{
		btl:   b,
		state: &InDraw{},
	}
}

// Exec executes current state
func (c *Context) Exec() error {
	return c.state.Exec(c)
}

// SetState changes the current state
func (c *Context) SetState(s IState) {
	c.state = s
}

// GetState returns the current state
func (c Context) GetState() IState {
	return c.state
}
