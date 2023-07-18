package bus

type IListener interface {
	Handler(event interface{})
}
