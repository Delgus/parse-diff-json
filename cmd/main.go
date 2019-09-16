package main

import (
	"fmt"
	"net/http"

	//api "github.com/delgus/parse-diff-json/api/v1"
	api "github.com/delgus/parse-diff-json/api/v2"
)

func main() {
	apiServer := api.NewApi()
	apiServer.OnMessage = func(message api.Message) {
		fmt.Println(message)
	}
	apiServer.OnEvent = func(event api.Event) {
		fmt.Println(event)
	}
	if err := http.ListenAndServe(":5000", apiServer); err != nil {
		panic(err)
	}
}
