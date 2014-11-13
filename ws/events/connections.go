package events

// eventConnections tracks a set of clients.
type eventConnections struct {
	connections map[*Client]bool // The set of know connections.
}

// NewEventConnections creates a new eventConections to track a set of clients.
func NewEventConnections() *eventConnections {
	return &eventConnections{
		connections: make(map[*Client]bool),
	}
}

// Register adds a new client to the list of known connections.
func (c *eventConnections) Register(client *Client) {
	c.connections[client] = true
}

// Unregister removes a client from the list of known connections. If
// the client doesn't exist it is ignored.
func (c *eventConnections) Unregister(client *Client) {
	if _, ok := c.connections[client]; ok {
		delete(c.connections, client)
	}
}

// Broadcast runs through the list of connections and sends the
// given message to them. If a send fails it assumes the client
// exited and it removes it from the list.
func (c *eventConnections) Broadcast(msg Message) {
	for conn := range c.connections {
		if err := conn.Send(msg); err != nil {
			// Bad send, assume connection is dead and delete
			delete(c.connections, conn)
		}
	}
}
