package match

// InResult is expressed the state of result that is implemented battle.IState
type InResult struct{}

// Exec checks the state of the result ending and changes the state
func (s InResult) Exec(ctx *Context) error {
	ctx.SetState(nil)
	return nil
}
