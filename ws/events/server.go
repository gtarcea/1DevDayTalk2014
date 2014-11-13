package events

import "code.google.com/p/go.net/websocket"

// Server implements a hub for handling multiple
// websocket connections.
type Server struct {
	hub Hub
}

// NewServer creates a new server.
func NewServer(hub Hub) *Server {
	return &Server{
		hub: hub,
	}
}

// OnConnection is called when a new websocket connection is made.
// It creates a persistent client connection and registers that
// connection with the hub. It it meant to be called by the
// websocket.Handler method.
func (s *Server) OnConnection(ws *websocket.Conn) {
	defer ws.Close()

	client := NewClient(ws, s.hub)
	s.hub.Register(client)
	client.Listen()
}
