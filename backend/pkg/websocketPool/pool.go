package websocketPool

// Pool broadcasts messages to the
// Clients and responsible for Clients being Registered.
type Pool struct {
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	UnRegister chan *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

func (pool *Pool) Run() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

		case client := <-pool.UnRegister:
			if _, ok := pool.Clients[client]; ok {
				delete(pool.Clients, client)
				close(client.send)
			}

		case message := <-pool.Broadcast:
			for client := range pool.Clients {
				select {
				case client.send <- message:
				default:
					delete(pool.Clients, client)
					close(client.send)
				}

			}
		}
	}
}
