package events

// eventHub implements a Hub for broadcasting messages to
// a connected set of clients.
type eventHub struct {
	connections EventConnections // Track the connections
	register    chan *Client     // Requests to register a new client
	unregister  chan *Client     // Requests to unregister (remove) an existing client
	broadcast   chan Message     // The channel to send broadcast messages on
}

// NewHub creates a new hub instance. It doesn't start it.
func NewHub(c EventConnections) *eventHub {
	return &eventHub{
		connections: c,
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		broadcast:   make(chan Message),
	}
}

// Start starts a go routine to process events.
func (h *eventHub) Start() {
	go h.processEvents()
}

// Register will register a new client.
func (h *eventHub) Register(c *Client) {
	h.register <- c
}

// Unregister will remove a client from the list of connections.
func (h *eventHub) Unregister(c *Client) {
	h.unregister <- c
}

// Broadcast will send the given message out to all known clients.
func (h *eventHub) Broadcast(msg Message) {
	h.broadcast <- msg
}

// processEvents is the heart of the hub. It waits on multiple channels
// to register, unregister or broadcast messages.
func (h *eventHub) processEvents() {
	for {
		select {
		case connection := <-h.register:
			h.connections.Register(connection)
		case connection := <-h.unregister:
			h.connections.Unregister(connection)
		case msg := <-h.broadcast:
			h.connections.Broadcast(msg)
		}
	}
}
