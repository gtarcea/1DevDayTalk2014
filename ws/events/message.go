package events

import "github.com/gtarcea/1DevDayTalk2014/schema"

type Message struct {
	Event string      `json:"event"`
	Data  schema.User `json:"data"`
}
