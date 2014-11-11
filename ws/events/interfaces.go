package events

type EventConnections interface {
	Register(*Client)
	Unregister(*Client)
	Broadcast(Message)
}

type Hub interface {
	Register(*Client)
	Unregister(*Client)
	Broadcast(Message)
	Start() error
}
