package events

type eventConnections struct {
	connections map[*Client]bool
}

func NewEventConnections() *eventConnections {
	return &eventConnections{
		connections: make(map[*Client]bool),
	}
}

func (c *eventConnections) Register(client *Client) {
	c.connections[client] = true
}

func (c *eventConnections) Unregister(client *Client) {
	if _, ok := c.connections[client]; ok {
		c.deleteConnection(client)
	}
}

func (c *eventConnections) deleteConnection(client *Client) {
	delete(c.connections, client)
}

func (c *eventConnections) Broadcast(msg Message) {
	for conn := range c.connections {
		if err := conn.Send(msg); err != nil {
			// Bad send, assume connection is dead and delete
			c.deleteConnection(conn)
		}
	}
}
