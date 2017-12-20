package match

type InResult struct{}

func (s InResult) Exec(ctx *Context) error {
	ctx.SetState(nil)
	return nil
}
