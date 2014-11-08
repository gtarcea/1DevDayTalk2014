package events

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type EventHandler struct {
	hub *Hub
}

type connection struct {
	ws   *websocket.Conn
	send chan interface{}
	hub  *Hub
}

func NewEventHandler(hub *Hub) *EventHandler {
	return &EventHandler{
		hub: hub,
	}
}

func (h *EventHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(resp, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		http.Error(resp, "Internal error", 500)
		return
	}

	connection := &connection{
		send: make(chan interface{}),
		ws:   ws,
		hub:  h.hub,
	}

	h.hub.register <- connection
	go connection.writeEvents()
	connection.readEvents()
}

func (c *connection) writeEvents() {

}

func (c *connection) readEvents() {

}
