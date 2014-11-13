package events

import (
	"io"

	"code.google.com/p/go.net/websocket"
)

// Create a buffered channel that can hold up to 10 items
const bufferedChannelSize = 10

// Client implements a persistent Websocket connection
type Client struct {
	ws   *websocket.Conn // websocket connection
	send chan Message    // channel to send requests to
	done chan bool       // channel to request client to stop
	hub  Hub             // Hub to broadcast on
}

// NewClient creates a new persistent websocket client
func NewClient(ws *websocket.Conn, hub Hub) *Client {
	return &Client{
		ws:   ws,
		send: make(chan Message, bufferedChannelSize),
		done: make(chan bool),
		hub:  hub,
	}
}

// Send will queue up a message to be sent on the websocket.
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

// Close will tell the client to exit and clean up the connection.
func (c *Client) Close() {
	c.done <- true
}

// Listen starts up a listener on the websocket, and a separate go routine
// for writing to the websocket.
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

// readListener processes messages on the websocket.
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
				// if msg.Event == "adduser" {
				// 	msg.Event = "addeduser"
				// 	c.hub.Broadcast(msg)
				// }
			}
		}
	}
}
