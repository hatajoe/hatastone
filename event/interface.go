package event

type EventNotify struct{}

type IEvent interface {
	GetID() EventID
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
