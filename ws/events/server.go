package events

import "code.google.com/p/go.net/websocket"

type Server struct {
	hub Hub
}

func NewServer(hub Hub) *Server {
	return &Server{
		hub: hub,
	}
}

func (s *Server) OnConnection(ws *websocket.Conn) {
	defer func() {
		ws.Close()
	}()
	client := NewClient(ws, s.hub)
	s.hub.Register(client)
	client.Listen()
}
