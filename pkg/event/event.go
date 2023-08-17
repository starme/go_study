package event

type Event struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
	//options *Options
}

func NewEvent(payload interface{}) Event {
	return Event{
		Name:    GetStructName(payload),
		Payload: payload,
		//options: &Options{},
	}
}
