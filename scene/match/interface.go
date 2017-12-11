package match

type IState interface {
	Exec(ctx *Context) error
}
