package events

type Connection interface {
	Send(interface{}) error
	Close() error
	Start()
}

type EventConnections interface {
	Register(Connection)
	Unregister(Connection)
	Broadcast(interface{})
}

type Hub interface {
	Register(Connection)
	Unregister(Connection)
	Broadcast(interface{})
	Start() error
}
