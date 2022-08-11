package main

import (
	"fmt"
	"net/http"
	"github.com/ssergomol/RealtimeChat/pkg/websocketPool"
)

func startPool() {
	pool := websocketPool.NewPool()
	go pool.Start()
}

func setRoutes() {
	http.HandleFunc("\ws", func(writer http.ResponseWriter, request *http.Request) {
		
	})
}

func main() {
	fmt.Println("Realtime Chat starting...")
	startPool()
	setRoutes()
	http.ListenAndServe(":8000", nil)
}
