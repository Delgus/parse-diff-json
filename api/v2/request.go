package v1

const (
	Unknown = iota
	MessageRequestType
	EventRequestType
)

type Request struct {
	Type int `json:"type"`
	Message
	Event
}

type Message struct {
	Text string `json:"text"`
}

type Event struct {
	Alert string `json:"alert"`
}
