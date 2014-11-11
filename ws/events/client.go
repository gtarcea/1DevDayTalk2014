package events

import (
	"io"

	"code.google.com/p/go.net/websocket"
)

const bufferedChannelSize = 10

type Client struct {
	ws   *websocket.Conn
	send chan Message
	done chan bool
	hub  Hub
}

func NewClient(ws *websocket.Conn, hub Hub) *Client {
	return &Client{
		ws:   ws,
		send: make(chan Message, bufferedChannelSize),
		done: make(chan bool),
		hub:  hub,
	}
}

func (c *Client) Send(msg Message) error {
	select {
	case c.send <- msg:
		return nil
	default:
		// No send channel means we have disconnected, time to
		// cleanup.
		return nil
	}
}

func (c *Client) Close() {
	c.done <- true
}

func (c *Client) Listen() {
	go c.writeListener()
	c.readListener()
}

// writeListener listens on the send channel for messages to write to
// the websocket.
func (c *Client) writeListener() {
	for {
		select {
		case msg := <-c.send:
			websocket.JSON.Send(c.ws, msg)
		case <-c.done:
			// Delete our connection
			c.hub.Unregister(c)

			// Tell reader to stop
			c.done <- true

			// exit process loop
			return
		}
	}
}

func (c *Client) readListener() {
	for {
		select {
		case <-c.done:
			// Delete our connection
			c.hub.Unregister(c)

			// Tell writer to stop
			c.done <- true

			//exit processing loop
			return
		default:
			// Read from websocket
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)

			switch {
			case err == io.EOF:
				// website connection closed
				c.done <- true
			case err != nil:
				// Handle error in some fashion
			default:
				// Do something with message
			}
			msg.Event = "addeduser"
			c.hub.Broadcast(msg)
		}
	}
}
