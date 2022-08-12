package websocketPool

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
	pool *Pool
	send chan []byte
}
