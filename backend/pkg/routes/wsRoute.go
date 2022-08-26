package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/websocketPool"
)

func RegisterWsRoute(router *mux.Router, pool *websocketPool.Pool) {
	router.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocketPool.ServeWS(pool, writer, request)
	})
}
