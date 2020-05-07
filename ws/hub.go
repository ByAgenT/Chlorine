package ws

// Hub is an aggregator of websocket connections as well as a provider of actions to ws clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Action dispatcher for clients.
	dispatcher *Dispatcher
}

// CreateHub creates blank instance of a ws Hub.
func CreateHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		dispatcher: &Dispatcher{},
	}
}

// Run start listening to register and unregister channels to accept and remove clients.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		}

	}
}

// AttachDispatcher register dispatcher to a Hub.
func (h *Hub) AttachDispatcher(dispatcher *Dispatcher) {
	h.dispatcher = dispatcher
}
