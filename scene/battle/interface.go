package battle

type IState interface {
	Exec(ctx *Context) error
}
