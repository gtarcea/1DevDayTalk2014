package events

type Hub struct {
	connections map[*connection]bool
	broadcast   chan interface{}
	register    chan *connection
	unregister  chan *connection
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[*connection]bool),
		broadcast:   make(chan interface{}),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
	}
}

func (h *Hub) Start() {
	go h.processEvents()
}

func (h *Hub) processEvents() {
	for {
		select {
		case connection := <-h.register:
			h.connections[connection] = true
		case connection := <-h.unregister:
			if h.connectionExists(connection) {
				h.deleteConnection(connection)
			}
		case msg := <-h.broadcast:
			h.broadcastMessage(msg)
		}
	}
}

func (h *Hub) connectionExists(connection *connection) bool {
	_, ok := h.connections[connection]
	return ok
}

func (h *Hub) deleteConnection(connection *connection) {
	delete(h.connections, connection)
	close(connection.send)
}

func (h *Hub) broadcastMessage(msg interface{}) {
	for connection := range h.connections {
		select {
		case connection.send <- msg:
		default:
			h.deleteConnection(connection)
		}
	}
}
