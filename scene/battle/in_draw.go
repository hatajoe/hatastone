package battle

// InDraw is expressed the state of draw the cards when beginning of the match that is implemented battle.IState
type InDraw struct{}

// Exec checks the state of draw the cards ending and changes the state
func (s InDraw) Exec(ctx *Context) error {
	ctx.SetState(&InMarigan{})
	return nil
}
