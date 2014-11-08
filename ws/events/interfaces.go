package events

type Connection interface {
	Key() interface{}
	Send(interface{}) error
	Close()
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
