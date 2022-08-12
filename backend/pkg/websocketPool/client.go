package websocketPool

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client is intermediary between websocket and Pool
type Client struct {
	Conn *websocket.Conn
	Pool *Pool
	Send chan Message
}

type Message struct {
	MessageType int    `json:"type"`
	Body        string `json:"body"`
}

// readPump pumps messages, which arrived from web client
// via webscocket connection, to the Pool

func (client *Client) readPump() {
	defer func() {
		client.Pool.unregister <- client
		client.Conn.Close()
	}()

	for {
		messageType, messageContent, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message := Message{MessageType: messageType, Body: string(messageContent)}
		client.Pool.broadcast <- message
	}
}

// writePump pumps messages, which were broadcasted on the Pool,
// back to the web client via websocket connection

func (client *Client) writePump() {
	defer func() {
		client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := client.Conn.WriteJSON(message); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func ServeWS(pool *Pool, writer http.ResponseWriter, request *http.Request) {
	conn, err := websocketPool.UpgradeHandler(writer, request)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(writer, err)
		}

		client := &websocketPool.Client{Conn: conn, Pool: pool, Send: make(chan websocketPool.Message)}
		client.Pool.
}
