package main

import (
	"fmt"
	"net/http"

	"github.com/ssergomol/RealtimeChat/pkg/websocketPool"
)

func main() {
	fmt.Println("Realtime Chat starting...")
	pool := websocketPool.NewPool()
	go pool.Run()

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocketPool.ServeWS(pool, writer, request)
	})

	http.ListenAndServe(":9000", nil)
}
