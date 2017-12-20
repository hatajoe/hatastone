package event

type Draw struct {
	ch EventChannel
}

func GetDrawEventID() EventID {
	return EventID("drawEvent")
}

func (e Draw) GetID() EventID {
	return GetDrawEventID()
}

func (e Draw) GetChannel() EventChannel {
	return e.ch
}

func (e *Draw) Open() {
	e.ch = make(EventChannel)
}

func (e *Draw) Close() {
	close(e.ch)
}

func (e Draw) Emit() {
	e.ch <- struct{}{}
}
