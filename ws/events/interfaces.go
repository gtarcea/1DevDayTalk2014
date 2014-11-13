package events

// EventConnections tracks a set of client connections
// and can be asked to broadcast a message to them.
type EventConnections interface {
	Register(*Client)
	Unregister(*Client)
	Broadcast(Message)
}

// Hub handles multiple connections, tracking and
// broadcasting to them.
type Hub interface {
	Register(*Client)
	Unregister(*Client)
	Broadcast(Message)
	Start()
}
