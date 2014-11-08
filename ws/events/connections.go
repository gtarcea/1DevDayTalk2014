package events

type eventConnections struct {
	connections map[Connection]bool
}

func newEventConnections() *eventConnections {
	return &eventConnections{
		connections: make(map[Connection]bool),
	}
}

func (c *eventConnections) Register(conn Connection) {
	c.connections[conn] = true
}

func (c *eventConnections) Unregister(conn Connection) {
	if _, ok := c.connections[conn]; ok {
		c.deleteConnection(conn)
	}
}

func (c *eventConnections) deleteConnection(conn Connection) {
	delete(c.connections, conn)
	conn.Close()
}
