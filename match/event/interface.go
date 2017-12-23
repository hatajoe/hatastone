package event

type EventID string

type IEventNotify interface{}

type IEvent interface {
	GetID() EventID
	Emit(IEventNotify) error
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
