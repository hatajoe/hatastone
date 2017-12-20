package event

type EventChannel chan struct{}

type IEvent interface {
	GetID() EventID
	GetChannel() EventChannel
	Open()
	Close()
	Emit()
}

type Events []IEvent

func (s Events) FindByID(id EventID) IEvent {
	for _, e := range s {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}
