package battle

// InMatch is expressed the state of match that is implemented battle.IState
type InMatch struct{}

// Exec checks the state of the match ending and changes the state
func (s InMatch) Exec(ctx *Context) error {
	ctx.SetState(&InResult{})
	return nil
}
