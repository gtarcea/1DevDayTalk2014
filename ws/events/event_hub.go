package events

type eventHub struct {
	connections EventConnections
	register    chan *Client
	unregister  chan *Client
	broadcast   chan Message
}

func NewHub(c EventConnections) *eventHub {
	return &eventHub{
		connections: c,
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		broadcast:   make(chan Message),
	}
}

func (h *eventHub) Start() error {
	go h.processEvents()
	return nil
}

func (h *eventHub) Register(c *Client) {
	h.register <- c
}

func (h *eventHub) Unregister(c *Client) {
	h.unregister <- c
}

func (h *eventHub) Broadcast(msg Message) {
	h.broadcast <- msg
}

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
