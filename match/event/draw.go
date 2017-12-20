package event

type Draw chan EventNotify

func GetDrawEventID() EventID {
	return EventID("drawEvent")
}

func (e Draw) GetID() EventID {
	return GetDrawEventID()
}

func (e Draw) Emit() {
	e <- EventNotify{}
}
