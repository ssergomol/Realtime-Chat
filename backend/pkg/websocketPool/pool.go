package websocketPool

import "fmt"

// Pool broadcasts messages to the
// clients and responsible for clients being Registered.
type Pool struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewPool() *Pool {
	return &Pool{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (pool *Pool) Run() {
	for {
		select {
		case client := <-pool.register:
			pool.clients[client] = true
			fmt.Println("New clinet registered!")

		case client := <-pool.unregister:
			if _, ok := pool.clients[client]; ok {
				delete(pool.clients, client)
				close(client.send)
				fmt.Println("Clinet unregistered!")
			}

		case message := <-pool.broadcast:
			fmt.Println("Find message from broadcast\nTrying to send for each client!")
			for client := range pool.clients {
				select {
				case client.send <- message:
					fmt.Println("Message sent to client on broadcast")
				default:
					delete(pool.clients, client)
					close(client.send)
				}

			}
		}
	}
}
