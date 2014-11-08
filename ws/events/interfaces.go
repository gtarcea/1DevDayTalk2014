package events

type Connection interface {
	Send(interface{}) error
	Close() error
}

type Connections interface {
	Register(Connection)
	Unregister(Connection)
	Broadcast(interface{})
}

type MessageHub interface {
	Register(Connection)
	Unregister(Connection)
	Broadcast(interface{})
}
