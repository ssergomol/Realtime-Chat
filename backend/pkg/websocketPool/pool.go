package websocketPool

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

		case client := <-pool.unregister:
			if _, ok := pool.clients[client]; ok {
				delete(pool.clients, client)
				close(client.send)
			}

		case message := <-pool.broadcast:
			for client := range pool.clients {
				select {
				case client.send <- message:
				default:
					delete(pool.clients, client)
					close(client.send)
				}

			}
		}
	}
}
