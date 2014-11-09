package events

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/gtarcea/1DevDayTalk2014/app"
)

type wsConnection struct {
	ws       *websocket.Conn
	send     chan interface{}
	upgrader websocket.Upgrader
}

func newWSConnection() *wsConnection {
	return
}

var defaultUpgrader = websocket.Upgrader{}

func open(resp http.ResponseWriter, req *http.Request, upgrader *websocket.Upgrader) (*wsConnection, error) {
	if req.Method != "GET" {
		return nil, app.ErrInvalid
	}
	u := defaultUpgrader
	if upgrader != nil {
		u = *upgrader
	}

	ws, err := u.Upgrade(resp, req, nil)
	if err != nil {
		return nil, err
	}
	conn := &wsConnection{
		ws:   ws,
		send: make(chan interface{}),
	}
	return conn, nil
}

func (c *wsConnection) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

}

func (c *wsConnection) Send(msg interface{}) error {
	c.send <- msg
	return nil
}

func (c *wsConnection) Close() error {
	return c.ws.Close()
}
