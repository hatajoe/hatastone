package match

type InMatch struct{}

func (s InMatch) Exec(ctx *Context) error {
	ctx.SetState(&InResult{})
	return nil
}
