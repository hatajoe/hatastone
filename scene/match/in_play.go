package match

type InPlay struct{}

func (s InPlay) Exec(ctx *Context) error {
	ctx.SetState(&InResult{})
	return nil
}
