package websocketPool

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client is intermediary between websocket and pool
type Client struct {
	conn *websocket.Conn
	pool *Pool
	send chan Message
}

type Message struct {
	MessageType  int    `json:"type"`
	Body         string `json:"body"`
	Notification bool   `json:"notification"`
}

// readPump pumps messages, which arrived from web client
// via webscocket connection, to the pool

func (client *Client) readPump() {
	defer func() {
		client.pool.unregister <- client
		client.conn.Close()
	}()

	for {
		MessageType, messageContent, err := client.conn.ReadMessage()
		fmt.Println("Got message from browser: ", string(messageContent))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message := Message{MessageType: MessageType, Body: string(messageContent), Notification: false}
		client.pool.broadcast <- message
	}
}

// writePump pumps messages, which were broadcasted on the pool,
// back to the web client via websocket connection

func (client *Client) writePump() {
	defer func() {
		client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:

			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := client.conn.WriteJSON(message); err != nil {
				log.Println(err)
				return
			}
			fmt.Println("Sent message to browser from Send channel!")
		}
	}
}

func ServeWS(pool *Pool, writer http.ResponseWriter, request *http.Request) {
	conn, err := UpgradeHandler(writer, request)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(writer, err)
	}
	fmt.Println("Protocol successfuly upgraded to WebScoket")

	client := &Client{conn: conn, pool: pool, send: make(chan Message)}
	client.pool.register <- client

	go client.readPump()
	go client.writePump()
}
