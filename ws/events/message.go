package events

import "github.com/gtarcea/1DevDayTalk2014/schema"

// Message is an event over the websocket. Right only
// one type of data is sent back.
type Message struct {
	Event string      `json:"event"`
	Data  schema.User `json:"data"`
}
