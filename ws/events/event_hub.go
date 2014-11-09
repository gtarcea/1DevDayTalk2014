package events

type eventHub struct {
	eventConnections EventConnections
	register         chan Connection
	unregister       chan Connection
	broadcast        chan interface{}
}

func NewHub(c EventConnections) *eventHub {
	return &eventHub{
		eventConnections: c,
		register:         make(chan Connection),
		unregister:       make(chan Connection),
		broadcast:        make(chan interface{}),
	}
}

func (h *eventHub) Start() error {
	go h.processEvents()
	return nil
}

func (h *eventHub) Register(conn Connection) {
	h.register <- conn
}

func (h *eventHub) Unregister(conn Connection) {
	h.unregister <- conn
}

func (h *eventHub) Broadcast(msg interface{}) {
	h.broadcast <- msg
}

func (h *eventHub) processEvents() {
	for {
		select {
		case connection := <-h.register:
			h.eventConnections.Register(connection)
		case connection := <-h.unregister:
			h.eventConnections.Unregister(connection)
		case msg := <-h.broadcast:
			h.eventConnections.Broadcast(msg)
		}
	}
}
