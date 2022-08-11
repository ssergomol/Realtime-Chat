package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssergomol/RealtimeChat/pkg/websocketPool"
)

func startPool() {
	pool := websocketPool.NewPool()
	go pool.Start()
}

func setRoutes() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := websocketPool.UpgradeHandler(writer, request)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(writer, err)
		}
	})
}

func main() {
	fmt.Println("Realtime Chat starting...")
	startPool()
	setRoutes()
	http.ListenAndServe(":8000", nil)
}
