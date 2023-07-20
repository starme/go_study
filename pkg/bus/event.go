package bus

type IEvent interface {
	Name() string
}

type BaseEvent struct{}
