package match

type InMarigan struct{}

func (s InMarigan) Exec(ctx *Context) error {
	ctx.SetState(&InPlay{})
	return nil
}
