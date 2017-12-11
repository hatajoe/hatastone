package match

import "github.com/hatajoe/hatastone/match"

// Context saves the state of match
type Context struct {
	btl   *match.Match
	state IState
}

// NewContext returns Context instance
func NewContext(b *match.Match) *Context {
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
