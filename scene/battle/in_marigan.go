package battle

// InMarigan is expressed the state of marigan that is implemented battle.IState
type InMarigan struct{}

// Exec checks the state of the marigan ending and changes the state
func (s InMarigan) Exec(ctx *Context) error {
	ctx.SetState(&InMatch{})
	return nil
}
