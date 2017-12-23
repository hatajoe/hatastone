package event

type notify struct {
	done Done
}

func (n notify) Done() {
	n.done.Emit(&DoneNotify{})
}
