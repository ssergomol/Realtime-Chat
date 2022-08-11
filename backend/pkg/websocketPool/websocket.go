package websocketPool

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func UpgradeHandler(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }

	return upgrader.Upgrade(writer, request, nil)
}
