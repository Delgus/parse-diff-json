package v1

import (
	"encoding/json"
	"net/http"
)

type Api struct {
	OnMessage func(message Message)
	OnEvent   func(event Event)
}

func NewApi() *Api {
	return &Api{
		OnMessage: func(message Message) {},
		OnEvent:   func(event Event) {},
	}
}

func (h *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req Request
	if err := decoder.Decode(&req); err != nil {
		panic(err)
	}
	switch req.Type {
	case MessageRequestType:
		h.OnMessage(req.Message)
	case EventRequestType:
		h.OnEvent(req.Event)
	default:
		panic("unknown type request")
	}
}
