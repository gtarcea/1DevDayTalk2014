package events

import "errors"

var ErrEventNotFound = errors.New("event not found")
var ErrCloseRequested = errors.New("close requested")

type HandlerFunc func(c *Client, data interface{}) error

type EventHandler struct {
	events map[string]HandlerFunc
}

func Handler() *EventHandler {
	return &EventHandler{
		events: make(map[string]HandlerFunc),
	}
}

func (h *EventHandler) On(eventName string, f HandlerFunc) *EventHandler {
	h.events[eventName] = f
	return h
}

func (h *EventHandler) handleEvent(c *Client, eventName string, data interface{}) error {
	if f, ok := h.events[eventName]; ok && f != nil {
		return f(c, data)
	}
	return ErrEventNotFound
}
