package main

import (
	"fmt"
	v1 "github.com/delgus/meetup/diffjson/api/v1"
	"net/http"
)

func main() {
	api := v1.NewApi()
	api.OnMessage = func(message v1.Message) {
		fmt.Println(message)
	}
	api.OnEvent = func(event v1.Event) {
		fmt.Println(event)
	}
	if err := http.ListenAndServe(":5000", api); err != nil {
		panic(err)
	}
}
