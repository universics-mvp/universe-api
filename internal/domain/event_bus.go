package event

type Event interface{}

type EventBus struct {
	messageHandlers map[string][]func(event Event) any
}

func NewEventBus() EventBus {
	return EventBus{
		messageHandlers: make(map[string][]func(event Event) any),
	}
}

func (bus *EventBus) On(name string, handler func(event Event) any) {
	if handlers, ok := bus.messageHandlers[name]; ok {
		bus.messageHandlers[name] = append(handlers, handler)
	} else {
		bus.messageHandlers[name] = []func(event Event) any{handler}
	}
}
func (bus EventBus) Emit(name string, event Event) {
	if handlers, ok := bus.messageHandlers[name]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
