package resolver

type Context struct {
	strategy IStrategy
}

func NewContext(st IStrategy) *Context {
	return &Context{strategy: st}
}

func (c Context) Resolve(inf IInfluencer, acc IAccepter) {
	c.strategy.Resolve(inf, acc)
}
